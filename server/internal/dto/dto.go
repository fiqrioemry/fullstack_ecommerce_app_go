package dto

import (
	"mime/multipart"
	"time"
)

// AUTHENTICATION  =================================
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Fullname string `json:"fullname" binding:"required,min=5"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SendOTPRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type VerifyOTPRequest struct {
	Email string `json:"email" binding:"required,email"`
	OTP   string `json:"otp" binding:"required,len=6"`
}

type AuthMeResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Avatar   string `json:"avatar"`
	Role     string `json:"role"`
}

type GoogleSignInRequest struct {
	IDToken string `json:"idToken" binding:"required"`
}

// AUTHENTICATION  =================================

// PROFILE & ADDRESS MANAGEMENT ====================

type ProfileResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Avatar   string `json:"avatar"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
	Phone    string `json:"phone"`
	JoinedAt string `json:"joinedAt"`
}

type UpdateProfileRequest struct {
	Fullname string `json:"fullname" binding:"required,min=5"`
	Birthday string `json:"birthday"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
}

type CreateAddressRequest struct {
	Name          string `json:"name" binding:"required"`
	Address       string `json:"address" binding:"required"`
	ProvinceID    uint   `json:"provinceId" binding:"required"`
	CityID        uint   `json:"cityId" binding:"required"`
	DistrictID    uint   `json:"districtId" binding:"required"`
	SubdistrictID uint   `json:"subdistrictId" binding:"required"`
	PostalCodeID  uint   `json:"postalCodeId" binding:"required"`
	Phone         string `json:"phone" binding:"required"`
	IsMain        bool   `json:"isMain"`
}

type AddressResponse struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	ProvinceID    int    `json:"provinceId"`
	Province      string `json:"province"`
	City          string `json:"city"`
	CityID        int    `json:"cityId"`
	District      string `json:"district"`
	DistrictID    int    `json:"districtId"`
	Subdistrict   string `json:"subdistrict"`
	SubdistrictID int    `json:"subdistrictId"`
	PostalCode    string `json:"postalCode"`
	PostalCodeID  int    `json:"postalCodeId"`
	Phone         string `json:"phone"`
	IsMain        bool   `json:"isMain"`
	CreatedAt     string `json:"createdAt"`
}

type AddressQueryParam struct {
	Q     string `form:"q"`
	Page  int    `form:"page"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}

type SearchCityRequest struct {
	Query string `json:"q" binding:"required"`
}

type SearchProvinceRequest struct {
	Query string `json:"q" binding:"required"`
}

type UpdateAddressRequest struct {
	Name          string `json:"name" binding:"required"`
	Address       string `json:"address" binding:"required"`
	ProvinceID    uint   `json:"provinceId" binding:"required"`
	CityID        uint   `json:"cityId" binding:"required"`
	DistrictID    uint   `json:"districtId" binding:"required"`
	SubdistrictID uint   `json:"subdistrictId" binding:"required"`
	PostalCodeID  uint   `json:"postalCodeId" binding:"required"`
	Phone         string `json:"phone" binding:"required"`
}

// PROFILE & ADDRESS MANAGEMENT ====================

// PRODUCT, CATEGORY, BANNER REQUEST & RESPONSE  =====================
type CategoryResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Image string `json:"image"`
}

type CategoryQueryParam struct {
	Page   int    `form:"page"`
	Limit  int    `form:"limit"`
	Search string `form:"q"`
	Sort   string `form:"sort"`
}

type CategoryListResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Image string `json:"image"`
}

type BannerRequest struct {
	Image    *multipart.FileHeader `form:"image" binding:"required"`
	ImageURL string                `form:"-"`
	Position string                `form:"position" binding:"required,oneof=top side-left side-right bottom"`
}

type BannerResponse struct {
	ID       string `json:"id"`
	Position string `json:"position"`
	Image    string `json:"image"`
}

type CreateCategoryRequest struct {
	Name     string                `form:"name" binding:"required,min=5"`
	Image    *multipart.FileHeader `form:"image" binding:"required"`
	ImageURL string                `form:"-"`
}

type UpdateCategoryRequest struct {
	Name     string                `form:"name" binding:"required,min=5"`
	Image    *multipart.FileHeader `form:"image" binding:"required"`
	ImageURL string                `form:"-"`
}

type CreateProductRequest struct {
	Name        string                  `form:"name" binding:"required,min=5"`
	CategoryID  string                  `form:"categoryId" binding:"required,uuid4"`
	Description string                  `form:"description" binding:"required,min=20"`
	Slug        string                  `form:"slug"`
	Price       float64                 `form:"price" binding:"required"`
	Stock       int                     `form:"stock" binding:"required"`
	Discount    *float64                `form:"discount"`
	IsActive    bool                    `form:"isActive"`
	IsFeatured  bool                    `form:"isFeatured"`
	Weight      float64                 `form:"weight" binding:"required"`
	Length      float64                 `form:"length" binding:"required"`
	Width       float64                 `form:"width" binding:"required"`
	Height      float64                 `form:"height" binding:"required"`
	Images      []*multipart.FileHeader `form:"images" binding:"omitempty"`
	ImageURLs   []string                `form:"-"`
}

type UpdateProductRequest struct {
	Name        string                  `form:"name" binding:"required,min=5"`
	CategoryID  string                  `form:"categoryId" binding:"required,uuid4"`
	Description string                  `form:"description" binding:"required,min=20"`
	Price       float64                 `form:"price" binding:"required"`
	Stock       int                     `form:"stock" binding:"required"`
	Discount    *float64                `form:"discount"`
	IsActive    bool                    `form:"isActive"`
	IsFeatured  bool                    `form:"isFeatured"`
	Weight      float64                 `form:"weight" binding:"required"`
	Length      float64                 `form:"length" binding:"required"`
	Width       float64                 `form:"width" binding:"required"`
	Height      float64                 `form:"height" binding:"required"`
	Images      []*multipart.FileHeader `form:"images" binding:"omitempty"`
	ImageURLs   []string                `form:"-"`
}

type ProductListResponse struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Slug          string   `json:"slug"`
	Discount      *float64 `json:"discount"`
	Description   string   `json:"description"`
	Price         float64  `json:"price"`
	CategoryID    string   `json:"categoryId"`
	AverageRating float64  `json:"averageRating"`
	Category      string   `json:"category"`
	IsActive      bool     `json:"isActive"`
	Height        float64  `json:"height"`
	Width         float64  `json:"width"`
	Length        float64  `json:"length"`
	Weight        float64  `json:"weight"`
	IsFeatured    bool     `json:"isFeatured"`
	Stock         int      `json:"stock"`
	Images        []string `json:"images"`
}

type PaginationResponse struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalRows  int `json:"totalRows"`
	TotalPages int `json:"totalPages"`
}

type GetAllProductsRequest struct {
	Search   string  `form:"q"`
	Status   string  `form:"status"`
	Category string  `form:"category"`
	MinPrice float64 `form:"minPrice"`
	MaxPrice float64 `form:"maxPrice"`
	Rating   float64 `form:"rating"`
	Sort     string  `form:"sort"`
	Page     int     `form:"page"`
	Limit    int     `form:"limit"`
}

type ProductDetailResponse struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Slug          string   `json:"slug"`
	Description   string   `json:"description"`
	Price         float64  `json:"price"`
	Stock         int      `json:"stock"`
	Discount      *float64 `json:"discount"`
	CategoryID    string   `json:"categoryId"`
	Category      string   `json:"category"`
	Height        float64  `json:"height"`
	Width         float64  `json:"width"`
	Length        float64  `json:"length"`
	Weight        float64  `json:"weight"`
	AverageRating float64  `json:"averageRating"`
	Images        []string `json:"images"`
}

// PRODUCT, CATEGORY, BANNER REQUEST & RESPONSE  =====================

// TRANSACTION REQUEST & RESPONSE  ================
type CartItemRequest struct {
	ProductID string `json:"productId" binding:"required,uuid4"`
	Quantity  int    `json:"quantity" binding:"required,min=1"`
}

type UpdateCartItemRequest struct {
	Quantity int `json:"quantity" binding:"required,min=1"`
}

type UnCheckedRequest struct {
	IsChecked bool `json:"isChecked" binding:"required"`
}

type CartItemResponse struct {
	ProductID        string  `json:"productId"`
	Name             string  `json:"name"`
	Price            float64 `json:"price"`
	Discount         float64 `json:"discount"`
	DiscountedPrice  float64 `json:"discountedPrice"`
	Image            string  `json:"image"`
	IsChecked        bool    `json:"isChecked"`
	Weight           float64 `json:"weight"`
	Quantity         int     `json:"quantity"`
	OriginalSubtotal float64 `json:"originalSubtotal"`
	Subtotal         float64 `json:"subtotal"`
}

type CartResponse struct {
	ID    string             `json:"id"`
	Items []CartItemResponse `json:"items"`
	Total float64            `json:"total"`
}

type CheckoutRequest struct {
	Courier      string  `json:"courier" binding:"required"`
	ShippingCost float64 `json:"shippingCost" binding:"required"`
	VoucherCode  *string `json:"voucherCode"`
	Note         *string `json:"note"`
}

type CheckoutResponse struct {
	PaymentID string `json:"paymentId"`
	SnapToken string `json:"snapToken"`
	SnapURL   string `json:"snapUrl"`
}

type CreateVoucherRequest struct {
	Code         string   `json:"code" binding:"required"`
	Description  string   `json:"description" binding:"required"`
	DiscountType string   `json:"discountType" binding:"required,oneof=fixed percentage"`
	Discount     float64  `json:"discount" binding:"required,gt=0"`
	MaxDiscount  *float64 `json:"maxDiscount,omitempty"`
	IsReusable   bool     `json:"isReusable"`
	Quota        int      `json:"quota" binding:"required,gt=0"`
	ExpiredAt    string   `json:"expiredAt" binding:"required,datetime=2006-01-02"`
}

type UpdateVoucherRequest struct {
	Description  string   `json:"description" binding:"required"`
	DiscountType string   `json:"discountType" binding:"required,oneof=fixed percentage"`
	Discount     float64  `json:"discount" binding:"required,gt=0"`
	MaxDiscount  *float64 `json:"maxDiscount,omitempty"`
	Quota        int      `json:"quota" binding:"required,gt=0"`
	IsReusable   bool     `json:"isReusable"`
	ExpiredAt    string   `json:"expiredAt" binding:"required,datetime=2006-01-02"`
}

type VoucherResponse struct {
	ID           string   `json:"id"`
	Code         string   `json:"code"`
	Description  string   `json:"description"`
	DiscountType string   `json:"discountType"`
	Discount     float64  `json:"discount"`
	MaxDiscount  *float64 `json:"maxDiscount,omitempty"`
	Quota        int      `json:"quota"`
	ExpiredAt    string   `json:"expiredAt"`
	CreatedAt    string   `json:"createdAt"`
}

type ApplyVoucherRequest struct {
	UserID *string `json:"userId"`
	Code   string  `json:"code" binding:"required"`
	Total  float64 `json:"total" binding:"required"`
}

type ApplyVoucherResponse struct {
	Code          string   `json:"code"`
	DiscountType  string   `json:"discountType"`
	Discount      float64  `json:"discount"`
	MaxDiscount   *float64 `json:"maxDiscount,omitempty"`
	DiscountValue float64  `json:"discountValue"`
	FinalTotal    float64  `json:"finalTotal"`
}

type PaymentResponse struct {
	ID            string  `json:"id"`
	UserID        string  `json:"userId"`
	InvoiceNumber string  `json:"invoiceNumber"`
	OrderID       string  `json:"orderID"`
	UserEmail     string  `json:"email"`
	Fullname      string  `json:"fullname"`
	Total         float64 `json:"total"`
	Method        string  `json:"method"`
	Status        string  `json:"status"`
	PaidAt        string  `json:"paidAt"`
}

type MidtransNotificationRequest struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}

type PaymentQueryParam struct {
	Page   int    `form:"page"`
	Limit  int    `form:"limit"`
	Search string `form:"q"`
	Sort   string `form:"sort"`
	Status string `form:"status"`
}

type OrderQueryParam struct {
	Page   int    `form:"page"`
	Limit  int    `form:"limit"`
	Search string `form:"q"`
	Sort   string `form:"sort"`   // e.g. "created_at desc"
	Status string `form:"status"` // optional filter
}

type OrderListResponse struct {
	ID            string          `json:"id"`
	UserID        string          `json:"userId"`
	InvoiceNumber string          `json:"invoiceNumber"`
	Items         []ItemsResponse `json:"items"`
	Status        string          `json:"status"`
	Total         float64         `json:"total"`
	PaymentLink   string          `json:"paymentLink"`
	CreatedAt     time.Time       `json:"createdAt"`
}

type ItemsResponse struct {
	ProductID   string `json:"id"`
	ProductName string `json:"name"`
	Quantity    int    `json:"quantity"`
	Image       string `json:"image"`
}

type OrderDetailResponse struct {
	ID              string  `json:"id"`
	InvoiceNumber   string  `json:"invoiceNumber"`
	TrackingCode    *string `json:"trackingCode"`
	CourierName     string  `json:"courierName"`
	UserID          string  `json:"userId"`
	RecipientName   string  `json:"recipientName"`
	Phone           string  `json:"phone"`
	ShippingCost    float64 `json:"shippingCost"`
	ShippingAddress string  `json:"shippingAddress"`
	Note            *string `json:"note"`
	Status          string  `json:"status"`
	Total           float64 `json:"total"`
	VoucherCode     *string `json:"voucherCode"`
	VoucherDiscount float64 `json:"voucherDiscount"`
	Tax             float64 `json:"tax"`

	AmountToPay float64               `json:"amountToPay"`
	CreatedAt   time.Time             `json:"createdAt"`
	Items       []ItemsDetailResponse `json:"items"`
}

type ItemsDetailResponse struct {
	ItemID      string  `json:"id"`
	ProductName string  `json:"name"`
	ProductSlug string  `json:"slug"`
	Image       string  `json:"image"`
	IsReviewed  bool    `json:"isReviewed"`
	Price       float64 `json:"price"`
	Discount    float64 `json:"discount"`
	Quantity    int     `json:"quantity"`
	Subtotal    float64 `json:"subtotal"`
}

type CreateShipmentRequest struct {
	TrackingCode string  `json:"trackingCode" binding:"required"`
	Notes        *string `json:"notes"`
}

type ShipmentResponse struct {
	OrderID      string     `json:"orderId"`
	TrackingCode string     `json:"trackingCode"`
	Status       string     `json:"status"`
	Notes        *string    `json:"notes,omitempty"`
	ShippedAt    *time.Time `json:"shippedAt,omitempty"`
	DeliveredAt  *time.Time `json:"deliveredAt,omitempty"`
}

type ConfirmDeliveryResponse struct {
	OrderID   string    `json:"orderId"`
	Status    string    `json:"status"`
	Delivered time.Time `json:"deliveredAt"`
}

type CreateReviewRequest struct {
	Rating   int                   `form:"rating" binding:"required,min=1,max=5"`
	Comment  string                `form:"comment" binding:"omitempty"`
	Image    *multipart.FileHeader `form:"image" binding:"required"`
	ImageURL string                `form:"-"`
}

type ReviewResponse struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	Fullname  string    `json:"fullname"`
	Avatar    string    `json:"avatar"`
	ProductID string    `json:"productId"`
	Rating    int       `json:"rating"`
	Comment   string    `json:"comment"`
	Image     *string   `json:"images,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}

type ShippingCostRequest struct {
	DestinationProvinceID int    `json:"provinceId" binding:"required"`
	DestinationCityID     int    `json:"cityId" binding:"required"`
	Weight                int    `json:"weight" binding:"required"`
	Courier               string `json:"courier" binding:"required"`
}

type CancelOrderResponse struct {
	OrderID string `json:"orderId"`
	Status  string `json:"status"`
}

// ORDER, TRANSACTION, REVIEW REQUEST & RESPONSE  ================

// NOTIFICATIONS REQUEST & RESPONSE ================
type NotificationSettingResponse struct {
	TypeID  string `json:"typeId"`
	Code    string `json:"code"`
	Title   string `json:"title"`
	Channel string `json:"channel"`
	Enabled bool   `json:"enabled"`
}

type UpdateNotificationSettingRequest struct {
	TypeID  string `json:"typeId" binding:"required"`
	Channel string `json:"channel" binding:"required,oneof=email browser"`
	Enabled bool   `json:"enabled"`
}

type CreateNotificationRequest struct {
	UserID   string `json:"userId"`
	TypeCode string `json:"typeCode"`
	Title    string `json:"title"`
	Message  string `json:"message"`
	Channel  string `json:"channel"`
}

type NotificationResponse struct {
	ID        string `json:"id"`
	TypeCode  string `json:"typeCode"`
	Title     string `json:"title"`
	Message   string `json:"message"`
	Channel   string `json:"channel"`
	IsRead    bool   `json:"isRead"`
	CreatedAt string `json:"createdAt"`
}

type SendNotificationRequest struct {
	TypeCode string `json:"typeCode" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Message  string `json:"message" binding:"required"`
}

// NOTIFICATIONS REQUEST & RESPONSE ================
type NotificationEvent struct {
	UserID  string `json:"userId"`
	Type    string `json:"type"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

// ADMIN DASHBOARD AND USER MANAGEMENT ==================
type CustomerQueryParam struct {
	Q     string `form:"q"`
	Sort  string `form:"sort"`
	Page  int    `form:"page,default=1"`
	Limit int    `form:"limit,default=10"`
}

type CustomerListResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Fullname  string `json:"fullname"`
	Phone     string `json:"phone"`
	Avatar    string `json:"avatar"`
	CreatedAt string `json:"createdAt"`
}

type CustomerDetailResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Fullname  string `json:"fullname"`
	Phone     string `json:"phone"`
	Avatar    string `json:"avatar"`
	Gender    string `json:"gender"`
	Birthday  string `json:"birthday,omitempty"`
	Address   string `json:"address"`
	CreatedAt string `json:"createdAt"`
	LastLogin string `json:"lastLogin,omitempty"`
}

type DashboardStatsResponse struct {
	TotalCustomers int64   `json:"totalCustomers"`
	TotalProducts  int64   `json:"totalProducts"`
	TotalOrders    int64   `json:"totalOrders"`
	TotalRevenue   float64 `json:"totalRevenue"`
}

type RevenueStatRequest struct {
	Range string `form:"range" binding:"omitempty,oneof=daily monthly yearly"`
}

type RevenueStat struct {
	Date  string  `json:"date"`
	Total float64 `json:"total"`
}

type RevenueStatsResponse struct {
	Range         string        `json:"range"`
	TotalRevenue  float64       `json:"totalRevenue"`
	RevenueSeries []RevenueStat `json:"revenueSeries"`
}
