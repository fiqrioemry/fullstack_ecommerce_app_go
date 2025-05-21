package services

import (
	"errors"
	"fmt"
	"log"
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"server/internal/utils"
	"time"

	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"gorm.io/gorm"
)

type OrderService interface {
	GetOrderDetail(orderID string) (*dto.OrderDetailResponse, error)
	GetShipmentByOrderID(orderID string) (*dto.ShipmentResponse, error)
	CancelOrder(orderID, userID string) (*dto.CancelOrderResponse, error)
	Checkout(userID string, req dto.CheckoutRequest) (*dto.CheckoutResponse, error)
	ConfirmOrderDelivered(orderID string, userID string) (*dto.ConfirmDeliveryResponse, error)
	CreateShipment(orderID string, req dto.CreateShipmentRequest) (*dto.ShipmentResponse, error)
	GetAllOrders(userID string, role string, param dto.OrderQueryParam) ([]dto.OrderListResponse, *dto.PaginationResponse, error)
}

type orderService struct {
	orderRepo           repositories.OrderRepository
	paymentRepo         repositories.PaymentRepository
	authRepo            repositories.AuthRepository
	productRepo         repositories.ProductRepository
	voucherService      VoucherService
	notificationService NotificationService
}

func NewOrderService(orderRepo repositories.OrderRepository, paymentRepo repositories.PaymentRepository, authRepo repositories.AuthRepository, productRepo repositories.ProductRepository, voucherService VoucherService, notificationService NotificationService) OrderService {
	return &orderService{orderRepo, paymentRepo, authRepo, productRepo, voucherService, notificationService}
}

func (s *orderService) Checkout(userID string, req dto.CheckoutRequest) (*dto.CheckoutResponse, error) {
	uid, _ := uuid.Parse(userID)

	user, err := s.authRepo.GetUserByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	carts, err := s.orderRepo.GetUserCart(uid)
	if err != nil || len(carts) == 0 {
		return nil, errors.New("cart is empty")
	}

	address, err := s.orderRepo.GetMainAddress(uid)
	if err != nil {
		return nil, errors.New("main address not found")
	}

	var total float64
	var items []models.OrderItem
	for _, c := range carts {
		if c.Quantity > c.Product.Stock {
			return nil, fmt.Errorf("stock not enough for product: %s", c.Product.Name)
		}

		image := ""
		if len(c.Product.ProductGallery) > 0 {
			image = c.Product.ProductGallery[0].Image
		}

		finalPrice := c.Product.Price
		if c.Product.Discount != nil && *c.Product.Discount > 0 {
			finalPrice -= *c.Product.Discount
			if finalPrice < 0 {
				finalPrice = 0
			}
		}
		items = append(items, models.OrderItem{
			ProductID:   c.Product.ID,
			ProductName: c.Product.Name,
			ProductSlug: c.Product.Slug,
			Image:       image,
			Price:       finalPrice,
			Quantity:    c.Quantity,
			Subtotal:    finalPrice * float64(c.Quantity),
		})
		total += finalPrice * float64(c.Quantity)
	}

	taxRate := utils.GetTaxRate()
	tax := total * taxRate

	var voucherDiscount float64
	if req.VoucherCode != nil {
		apply, err := s.voucherService.ApplyVoucher(userID, dto.ApplyVoucherRequest{
			Code:  *req.VoucherCode,
			Total: total,
		})
		if err == nil {
			total = apply.FinalTotal
			voucherDiscount = apply.DiscountValue
		}
	}

	amountToPay := total + req.ShippingCost + tax
	orderID := uuid.New()
	invoice := utils.GenerateInvoiceNumber(orderID)

	order := &models.Order{
		ID:              orderID,
		InvoiceNumber:   invoice,
		UserID:          uid,
		AddressID:       address.ID,
		Courier:         req.Courier,
		ShippingCost:    req.ShippingCost,
		Tax:             tax,
		Note:            req.Note,
		Total:           total,
		AmountToPay:     amountToPay,
		VoucherCode:     req.VoucherCode,
		VoucherDiscount: voucherDiscount,
		Status:          "waiting_payment",
	}

	if err := s.orderRepo.CreateOrder(order); err != nil {
		return nil, err
	}

	for i := range items {
		items[i].OrderID = order.ID
	}
	if err := s.orderRepo.CreateOrderItems(items); err != nil {
		return nil, err
	}

	for _, c := range carts {
		err := s.productRepo.DecreaseProductStock(c.ProductID, c.Quantity)
		if err != nil {
			return nil, fmt.Errorf("failed to decrease stock for %s", c.Product.Name)
		}
	}

	if err := s.orderRepo.ClearUserCart(uid); err != nil {
		return nil, err
	}

	paymentID := uuid.New()
	payment := models.Payment{
		ID:      paymentID,
		UserID:  uid,
		OrderID: orderID,
		Method:  "-",
		Status:  "pending",
		PaidAt:  time.Time{},
		Total:   amountToPay,
	}

	if err := s.paymentRepo.CreatePayment(&payment); err != nil {
		return nil, err
	}

	// Send notification : success create new order
	// TODO : Replace using rabbitMQ later for sending notification
	payload := dto.NotificationEvent{
		UserID: user.ID.String(),
		Type:   "pending_payment",
		Message: fmt.Sprintf("Thank you %s, your order with invoice no. %s is created. Please complete your payment",
			user.Profile.Fullname, invoice),
	}

	err = s.notificationService.SendToUser(payload)
	if err != nil {
		log.Printf("Failed to send to user %s: %v\n", payload.UserID, err)
	}

	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  paymentID.String(),
			GrossAmt: int64(amountToPay),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Profile.Fullname,
			Email: user.Email,
			Phone: user.Profile.Phone,
		},
	}

	snapResp, _ := config.SnapClient.CreateTransaction(snapReq)

	order.PaymentLink = snapResp.RedirectURL
	if err := s.orderRepo.UpdateOrder(order); err != nil {
		return nil, err
	}

	return &dto.CheckoutResponse{
		PaymentID: paymentID.String(),
		SnapToken: snapResp.Token,
		SnapURL:   snapResp.RedirectURL,
	}, nil
}

func (s *orderService) GetAllOrders(userID string, role string, param dto.OrderQueryParam) ([]dto.OrderListResponse, *dto.PaginationResponse, error) {
	var (
		orders []models.Order
		total  int64
		err    error
	)

	switch role {
	case "admin":
		orders, total, err = s.orderRepo.GetAllOrders(param)
	case "customer":
		orders, total, err = s.orderRepo.GetOrdersByUserID(userID, param)
	default:
		return nil, nil, errors.New("unauthorized role")
	}

	if err != nil {
		return nil, nil, err
	}

	var result []dto.OrderListResponse
	for _, o := range orders {
		var items []dto.ItemsResponse
		for _, i := range o.Items {
			items = append(items, dto.ItemsResponse{
				ProductID:   i.ProductID.String(),
				ProductName: i.ProductName,
				Image:       i.Image,
				Quantity:    i.Quantity,
			})
		}

		result = append(result, dto.OrderListResponse{
			ID:          o.ID.String(),
			UserID:      o.UserID.String(),
			Items:       items,
			Status:      o.Status,
			Total:       o.AmountToPay,
			PaymentLink: o.PaymentLink,
			CreatedAt:   o.CreatedAt,
		})
	}

	totalPages := int((total + int64(param.Limit) - 1) / int64(param.Limit))
	pagination := &dto.PaginationResponse{
		Page:       param.Page,
		Limit:      param.Limit,
		TotalRows:  int(total),
		TotalPages: totalPages,
	}

	return result, pagination, nil
}

func (s *orderService) GetOrderDetail(orderID string) (*dto.OrderDetailResponse, error) {
	order, err := s.orderRepo.GetOrderDetail(orderID)
	if err != nil {
		return nil, err
	}

	var items []dto.ItemsDetailResponse
	for _, i := range order.Items {
		items = append(items, dto.ItemsDetailResponse{
			ProductID:   i.ProductID.String(),
			ProductName: i.ProductName,
			ProductSlug: i.ProductSlug,
			Image:       i.Image,
			Price:       i.Price,
			Quantity:    i.Quantity,
			Subtotal:    i.Subtotal,
		})
	}

	return &dto.OrderDetailResponse{
		ID:              order.ID.String(),
		InvoiceNumber:   order.InvoiceNumber,
		CourierName:     order.Courier,
		ShipmentID:      order.Shipment.ID.String(),
		UserID:          order.UserID.String(),
		CustomerName:    order.Address.Name,
		Phone:           order.Address.Phone,
		Address:         order.Address.Address,
		Province:        order.Address.Province,
		City:            order.Address.City,
		District:        order.Address.District,
		Subdistrict:     order.Address.Subdistrict,
		PostalCode:      order.Address.PostalCode,
		Note:            order.Note,
		Status:          order.Status,
		Total:           order.Total,
		VoucherCode:     order.VoucherCode,
		VoucherDiscount: order.VoucherDiscount,
		Tax:             order.Tax,
		ShippingCost:    order.ShippingCost,
		AmountToPay:     order.AmountToPay,
		CreatedAt:       order.CreatedAt,
		Items:           items,
	}, nil

}

func (s *orderService) CreateShipment(orderID string, req dto.CreateShipmentRequest) (*dto.ShipmentResponse, error) {
	id, err := uuid.Parse(orderID)
	if err != nil {
		return nil, errors.New("invalid order ID")
	}

	order, err := s.orderRepo.GetOrderDetail(orderID)
	if err != nil {
		return nil, errors.New("order not found")
	}

	if order.Shipment.ID != uuid.Nil {
		return nil, errors.New("shipment already exists for this order")
	}

	now := time.Now()
	shipment := &models.Shipment{
		ID:           uuid.New(),
		OrderID:      id,
		TrackingCode: req.TrackingCode,
		Status:       "pending",
		Notes:        req.Notes,
		ShippedAt:    &now,
	}

	err = s.orderRepo.WithTx(func(tx *gorm.DB) error {
		if err := tx.Create(shipment).Error; err != nil {
			return err
		}
		return tx.Model(&models.Order{}).
			Where("id = ?", id).
			Update("status", "success").Error
	})
	if err != nil {
		return nil, err
	}
	// TODO: Replace with RabbitMQ for async notification dispatch ---
	// ? Send notification : success shipment info
	payload := dto.NotificationEvent{
		UserID:  order.UserID.String(),
		Type:    "order_shipped",
		Message: "Thank you for your payment. Your order is now being processed.",
	}

	err = s.notificationService.SendToUser(payload)
	if err != nil {
		log.Printf("Fail to send notification to user %s: %v\n", payload.UserID, err)
	}
	// TODO: Replace with RabbitMQ for async notification dispatch ---

	return &dto.ShipmentResponse{
		OrderID:      shipment.OrderID.String(),
		TrackingCode: shipment.TrackingCode,
		Status:       shipment.Status,
		Notes:        shipment.Notes,
		ShippedAt:    shipment.ShippedAt,
	}, nil
}

func (s *orderService) GetShipmentByOrderID(orderID string) (*dto.ShipmentResponse, error) {
	id, err := uuid.Parse(orderID)
	if err != nil {
		return nil, errors.New("invalid order ID")
	}

	shipment, err := s.orderRepo.GetShipmentByOrderID(id)
	if err != nil {
		return nil, err
	}

	return &dto.ShipmentResponse{
		OrderID:      shipment.OrderID.String(),
		TrackingCode: shipment.TrackingCode,
		Status:       shipment.Status,
		Notes:        shipment.Notes,
		ShippedAt:    shipment.ShippedAt,
		DeliveredAt:  shipment.DeliveredAt,
	}, nil
}

func (s *orderService) ConfirmOrderDelivered(orderID string, userID string) (*dto.ConfirmDeliveryResponse, error) {
	id, err := uuid.Parse(orderID)
	if err != nil {
		return nil, errors.New("invalid order ID")
	}

	order, err := s.orderRepo.GetOrderDetail(orderID)
	if err != nil {
		return nil, errors.New("order not found")
	}

	if order.UserID.String() != userID {
		return nil, errors.New("unauthorized")
	}

	if order.Shipment.Status == "delivered" {
		return nil, errors.New("order already marked as delivered")
	}

	err = s.orderRepo.MarkOrderDelivered(id)
	if err != nil {
		return nil, err
	}

	return &dto.ConfirmDeliveryResponse{
		OrderID:   orderID,
		Status:    "delivered",
		Delivered: time.Now(),
	}, nil
}

func (s *orderService) CancelOrder(orderID, userID string) (*dto.CancelOrderResponse, error) {
	order, err := s.orderRepo.GetOrderDetail(orderID)
	if err != nil {
		return nil, errors.New("order not found")
	}

	// Validasi user dan status
	if order.UserID.String() != userID {
		return nil, errors.New("unauthorized")
	}

	payment, err := s.paymentRepo.GetPaymentByOrderID(order.ID.String())
	if err != nil {
		return nil, errors.New("payment not found")
	}
	if payment.Status != "pending" {
		return nil, errors.New("order cannot be canceled because payment already processed")
	}

	// Update status payment & order
	payment.Status = "failed"
	if err := s.paymentRepo.UpdatePayment(payment); err != nil {
		return nil, err
	}

	if err := s.productRepo.RestoreStockOnPaymentFailure(order); err != nil {
		return nil, err
	}

	order.Status = "canceled"
	if err := s.orderRepo.UpdateOrder(order); err != nil {
		return nil, err
	}

	return &dto.CancelOrderResponse{
		OrderID: order.ID.String(),
		Status:  "canceled",
	}, nil
}
