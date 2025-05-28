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
	GetAllOrders(userID string, role string, param dto.OrderQueryParam) ([]dto.OrderListResponse, *dto.PaginationResponse, error)
	Checkout(userID string, req dto.CheckoutRequest) (*dto.CheckoutResponse, error)
	GetOrderDetail(orderID string) (*dto.OrderDetailResponse, error)
	CreateShipment(orderID string, req dto.CreateShipmentRequest) (*dto.ShipmentResponse, error)
	GetShipmentByOrderID(orderID string) (*dto.ShipmentResponse, error)
	ConfirmOrderDelivered(orderID string) (*dto.ConfirmDeliveryResponse, error)
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

		price := c.Product.Price

		if c.Product.Discount != nil && *c.Product.Discount > 0 {
			log.Printf("calculating discount for product %s", c.Product.Name)
			discount := (price * *c.Product.Discount) / 100
			price -= discount
			if price < 0 {
				price = 0
			}
		}

		subtotal := price * float64(c.Quantity)
		total += subtotal

		items = append(items, models.OrderItem{
			ProductID:   c.Product.ID,
			ProductName: c.Product.Name,
			ProductSlug: c.Product.Slug,
			Image:       image,
			Price:       price,
			Discount:    c.Product.Discount,
			Quantity:    c.Quantity,
			Subtotal:    subtotal,
		})
	}

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

	taxRate := utils.GetTaxRate()
	tax := total * taxRate
	amountToPay := total + req.ShippingCost + tax

	orderID := uuid.New()
	invoice := utils.GenerateInvoiceNumber(orderID)

	order := &models.Order{
		ID:              orderID,
		InvoiceNumber:   invoice,
		UserID:          uid,
		Courier:         req.Courier,
		RecipientName:   user.Profile.Fullname,
		ShippingCost:    req.ShippingCost,
		ShippingAddress: fmt.Sprintf("%s, %s, %s, %s, %s", address.Address, address.Province, address.City, address.District, address.PostalCode),
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
		if err := s.productRepo.DecreaseProductStock(c.ProductID, c.Quantity); err != nil {
			return nil, fmt.Errorf("failed to decrease stock for %s", c.Product.Name)
		}
	}

	if err := s.orderRepo.ClearUserCart(uid); err != nil {
		return nil, err
	}

	paymentID := uuid.New()
	payment := models.Payment{
		ID:       paymentID,
		UserID:   uid,
		Fullname: user.Profile.Fullname,
		Email:    user.Email,
		OrderID:  orderID,
		Method:   "midtrans",
		Status:   "pending",
		PaidAt:   time.Time{},
		Total:    amountToPay,
	}

	if err := s.paymentRepo.CreatePayment(&payment); err != nil {
		return nil, err
	}

	if order.VoucherCode != nil && *order.VoucherCode != "" {
		if err := s.voucherService.DecreaseQuota(order.UserID, *order.VoucherCode); err != nil {
			return nil, fmt.Errorf("failed to decrease voucher quota")
		}
	}

	var itemDetails []midtrans.ItemDetails
	for _, item := range items {
		name := item.ProductName
		if len(name) > 45 {
			name = name[:45]
		}

		itemDetails = append(itemDetails, midtrans.ItemDetails{
			ID:    item.ProductID.String(),
			Name:  name,
			Price: int64(item.Price),
			Qty:   int32(item.Quantity),
		})
	}

	if req.ShippingCost > 0 {
		itemDetails = append(itemDetails, midtrans.ItemDetails{
			ID:    "shipping",
			Name:  fmt.Sprintf("Shipping via %s", order.Courier),
			Price: int64(req.ShippingCost),
			Qty:   1,
		})
	}
	if int64(voucherDiscount) > 0 {
		itemDetails = append(itemDetails, midtrans.ItemDetails{
			ID:    "discount",
			Name:  "Voucher Discount",
			Price: -int64(voucherDiscount),
			Qty:   1,
		})
	}

	if tax > 0 {
		itemDetails = append(itemDetails, midtrans.ItemDetails{
			ID:    "tax",
			Name:  "Tax (PPN)",
			Price: int64(tax),
			Qty:   1,
		})
	}

	snapRequest := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  payment.OrderID.String(),
			GrossAmt: int64(amountToPay),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Profile.Fullname,
			Email: user.Email,
			Phone: address.Phone,
		},
		Items:        &itemDetails,
		CustomField1: fmt.Sprintf("%s, %s, %s", address.Address, address.City, address.PostalCode),
		EnabledPayments: []snap.SnapPaymentType{
			snap.PaymentTypeGopay,
			snap.PaymentTypeBankTransfer,
			snap.PaymentTypeCreditCard,
		},
	}

	var sum int64
	for _, item := range itemDetails {
		sum += item.Price * int64(item.Qty)
	}
	log.Printf("✅ gross_amount = %d | sum(itemDetails) = %d", int64(amountToPay), sum)

	snapResp, err := config.SnapClient.CreateTransaction(snapRequest)

	order.PaymentLink = snapResp.RedirectURL
	if err := s.orderRepo.UpdateOrder(order); err != nil {
		return nil, err
	}

	err = s.notificationService.SendToUser(dto.NotificationEvent{
		UserID: user.ID.String(),
		Type:   "pending_payment",
		Title:  "Order Created",
		Message: fmt.Sprintf("Thank you %s, your order with invoice no. %s is created. Please complete your payment.",
			user.Profile.Fullname, invoice),
	})
	if err != nil {
		log.Printf("Fail to send notification to user %s: %v\n", user.ID.String(), err)
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
			ID:            o.ID.String(),
			UserID:        o.UserID.String(),
			InvoiceNumber: o.InvoiceNumber,
			Items:         items,
			Status:        o.Status,
			Total:         o.AmountToPay,
			PaymentLink:   o.PaymentLink,
			CreatedAt:     o.CreatedAt,
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
			ItemID:      i.ID.String(),
			ProductName: i.ProductName,
			ProductSlug: i.ProductSlug,
			Image:       i.Image,
			Price:       i.Price,
			IsReviewed:  i.IsReviewed,
			Quantity:    i.Quantity,
			Subtotal:    i.Subtotal,
		})
	}

	return &dto.OrderDetailResponse{
		ID:              order.ID.String(),
		InvoiceNumber:   order.InvoiceNumber,
		TrackingCode:    &order.Shipment.TrackingCode,
		CourierName:     order.Courier,
		UserID:          order.UserID.String(),
		RecipientName:   order.RecipientName,
		ShippingAddress: order.ShippingAddress,
		ShippingCost:    order.ShippingCost,
		Phone:           order.Phone,
		Note:            order.Note,
		Status:          order.Status,
		Total:           order.Total,
		VoucherCode:     order.VoucherCode,
		VoucherDiscount: order.VoucherDiscount,
		Tax:             order.Tax,
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

	now := time.Now()
	shipment := &models.Shipment{
		ID:           uuid.New(),
		OrderID:      id,
		TrackingCode: req.TrackingCode,
		Status:       "shipped",
		Notes:        req.Notes,
		ShippedAt:    &now,
	}

	err = s.orderRepo.WithTx(func(tx *gorm.DB) error {
		if err := tx.Create(shipment).Error; err != nil {
			return err
		}
		return tx.Model(&models.Order{}).
			Where("id = ?", id).
			Update("status", "process").Error
	})
	if err != nil {
		return nil, err
	}

	// ? Waiting for order is shipped notifications : event 3
	// TODO: Replace with RabbitMQ for async notification dispatch ---
	// ? Send notification : success shipment info
	payload := dto.NotificationEvent{
		UserID:  order.UserID.String(),
		Type:    "order_shipped",
		Message: "Your Order is being shipped and on way to your destination",
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

func (s *orderService) ConfirmOrderDelivered(orderID string) (*dto.ConfirmDeliveryResponse, error) {
	id, err := uuid.Parse(orderID)
	if err != nil {
		return nil, errors.New("invalid order ID")
	}

	order, err := s.orderRepo.GetOrderDetail(orderID)
	if err != nil {
		return nil, errors.New("order not found")
	}

	if order.Shipment.Status == "delivered" {
		return nil, errors.New("order already marked as delivered")
	}

	err = s.orderRepo.MarkOrderDelivered(id)
	if err != nil {
		return nil, err
	}

	// ? Waiting for order is completed notifications : event 4
	// TODO: Replace with RabbitMQ for async notification dispatch ---
	payload := dto.NotificationEvent{
		UserID:  order.UserID.String(),
		Title:   "Order Delivered",
		Type:    "order_completed",
		Message: fmt.Sprintf("Your order #%s has been successfully delivered. Thank you for shopping with us!", order.InvoiceNumber),
	}

	err = s.notificationService.SendToUser(payload)
	if err != nil {
		log.Printf("Fail to send notification to user %s: %v\n", payload.UserID, err)
	}
	// TODO: Replace with RabbitMQ for async notification dispatch ---

	err = s.orderRepo.WithTx(func(tx *gorm.DB) error {
		return tx.Model(&models.Order{}).
			Where("id = ?", id).
			Update("status", "success").Error
	})
	if err != nil {
		return nil, err
	}

	return &dto.ConfirmDeliveryResponse{
		OrderID:   orderID,
		Status:    "delivered",
		Delivered: time.Now(),
	}, nil
}
