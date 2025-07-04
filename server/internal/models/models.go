package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// USER SERVICES MODEL ================================
type User struct {
	ID        uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	Email     string         `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password  string         `gorm:"type:text;not null" json:"-"`
	Role      string         `gorm:"type:varchar(255);default:'customer';check:role IN ('customer','admin')" json:"role"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Profile   Profile   `gorm:"foreignKey:UserID" json:"profile"`
	Tokens    []Token   `gorm:"foreignKey:UserID" json:"-"`
	Addresses []Address `gorm:"foreignKey:UserID" json:"addresses,omitempty"`
}

type Token struct {
	ID        uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	UserID    uuid.UUID      `gorm:"type:char(36);index;not null" json:"userId"`
	Token     string         `gorm:"type:text;not null" json:"token"`
	ExpiredAt time.Time      `json:"expiredAt"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Profile struct {
	ID        uuid.UUID  `gorm:"type:char(36);primaryKey" json:"id"`
	UserID    uuid.UUID  `gorm:"type:char(36);uniqueIndex;not null" json:"userId"`
	Fullname  string     `gorm:"type:varchar(255);not null" json:"fullname"`
	Birthday  *time.Time `json:"birthday,omitempty"`
	Phone     string     `gorm:"type:varchar(20)" json:"phone"`
	Gender    string     `gorm:"type:varchar(10)" json:"gender"`
	Avatar    string     `gorm:"type:varchar(255)" json:"avatar"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
}

type Address struct {
	ID            uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	UserID        uuid.UUID      `gorm:"type:char(36);not null;index" json:"-"`
	Name          string         `gorm:"type:varchar(255);not null" json:"name"`
	IsMain        bool           `gorm:"default:true" json:"isMain"`
	Address       string         `gorm:"type:text;not null" json:"address"`
	ProvinceID    uint           `gorm:"not null" json:"provinceId"`
	CityID        uint           `gorm:"not null" json:"cityId"`
	DistrictID    uint           `gorm:"not null" json:"districtId"`
	SubdistrictID uint           `gorm:"not null" json:"subdistrictId"`
	PostalCodeID  uint           `gorm:"not null" json:"postalCodeId"`
	Province      string         `gorm:"type:varchar(255);not null" json:"province"`
	City          string         `gorm:"type:varchar(255);not null" json:"city"`
	District      string         `gorm:"type:varchar(255);not null" json:"district"`
	Subdistrict   string         `gorm:"type:varchar(255);not null" json:"subdistrict"`
	PostalCode    string         `gorm:"type:varchar(20);not null" json:"postalCode"`
	Phone         string         `gorm:"type:varchar(20);not null" json:"phone"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

type Province struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"type:varchar(100);not null" json:"name"`
	Cities []City `gorm:"foreignKey:ProvinceID" json:"-"`
}

type City struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	ProvinceID uint       `gorm:"not null" json:"provinceId"`
	Name       string     `gorm:"type:varchar(100);not null" json:"name"`
	Districts  []District `gorm:"foreignKey:CityID" json:"-"`
}

type District struct {
	ID           uint          `gorm:"primaryKey" json:"id"`
	CityID       uint          `gorm:"not null" json:"cityId"`
	Name         string        `gorm:"type:varchar(100);not null" json:"name"`
	Subdistricts []Subdistrict `gorm:"foreignKey:DistrictID" json:"-"`
}

type Subdistrict struct {
	ID          uint         `gorm:"primaryKey" json:"id"`
	DistrictID  uint         `gorm:"not null" json:"districtId"`
	Name        string       `gorm:"type:varchar(100);not null" json:"name"`
	PostalCodes []PostalCode `gorm:"foreignKey:SubdistrictID" json:"-"`
}

type PostalCode struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	ProvinceID    uint   `gorm:"not null" json:"provinceId"`
	CityID        uint   `gorm:"not null" json:"cityId"`
	DistrictID    uint   `gorm:"not null" json:"districtId"`
	SubdistrictID uint   `gorm:"not null" json:"subdistrictId"`
	PostalCode    string `gorm:"type:varchar(20);not null" json:"postalCode"`
}

// USER SERVICES MODEL ================================

// PRODUCT SERVICES MODEL ================================
type Category struct {
	ID        uuid.UUID      `gorm:"type:char(36);primaryKey"`
	Name      string         `gorm:"type:varchar(100);not null;unique" json:"name"`
	Slug      string         `gorm:"type:varchar(100);uniqueIndex" json:"slug"`
	Image     string         `gorm:"type:varchar(255)" json:"image"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Product struct {
	ID            uuid.UUID      `gorm:"type:char(36);primaryKey"`
	CategoryID    uuid.UUID      `gorm:"type:char(36);not null"`
	Name          string         `gorm:"type:varchar(255);not null"`
	Slug          string         `gorm:"type:varchar(255);uniqueIndex"`
	Description   string         `gorm:"type:text"`
	Stock         int            `gorm:"default:0" json:"stock"`
	Sold          int            `gorm:"default:0"`
	Price         float64        `gorm:"type:decimal(10,2);default:0" json:"price"`
	IsFeatured    bool           `gorm:"default:false"`
	IsActive      bool           `gorm:"default:true"`
	AverageRating float64        `gorm:"type:decimal(3,2);default:0" json:"averageRating"`
	Weight        float64        `gorm:"default:1000" json:"weight"`
	Length        float64        `gorm:"default:0" json:"length"`
	Width         float64        `gorm:"default:0" json:"width"`
	Height        float64        `gorm:"default:0" json:"height"`
	Discount      *float64       `gorm:"type:decimal(10,2);default:0" json:"discount"`
	CreatedAt     time.Time      `gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`

	Category       Category         `gorm:"foreignKey:CategoryID"`
	Review         []Review         `gorm:"foreignKey:ProductID"`
	ProductGallery []ProductGallery `gorm:"foreignKey:ProductID"`
}

type ProductGallery struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	ProductID uuid.UUID `gorm:"type:char(36);not null"`
	Image     string    `gorm:"type:varchar(255)" json:"image"`
}

type Review struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:char(36);not null" json:"userId"`
	Image     *string   `gorm:"type:varchar(255)" json:"image"`
	ProductID uuid.UUID `gorm:"type:char(36);not null" json:"productId"`
	Rating    int       `gorm:"not null" json:"rating"`
	Comment   string    `gorm:"type:text" json:"comment"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	User User `gorm:"foreignKey:UserID" json:"user"`
}

type Banner struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	Position  string    `gorm:"type:varchar(50);not null"`
	Image     string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// PRODUCT SERVICES MODEL ================================

// TRANSACTION SERVICES MODEL ================================
type Cart struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	UserID    uuid.UUID `gorm:"type:char(36);index"`
	ProductID uuid.UUID `gorm:"type:char(36);not null"`
	Quantity  int       `gorm:"default:1"`
	IsChecked bool      `gorm:"default:true"`
	Product   Product   `gorm:"foreignKey:ProductID"`
}

type Order struct {
	ID              uuid.UUID      `gorm:"type:char(36);primaryKey"`
	UserID          uuid.UUID      `gorm:"type:char(36);not null;index"`
	PaymentLink     string         `gorm:"type:text"`
	InvoiceNumber   string         `gorm:"type:varchar(100);uniqueIndex;not null"`
	Note            *string        `gorm:"type:text"`
	Courier         string         `gorm:"type:varchar(100)"`
	Status          string         `gorm:"type:varchar(20);default:'waiting_payment';check:status IN ('waiting_payment','canceled', 'pending', 'process', 'success')" json:"status"`
	Total           float64        `gorm:"type:decimal(10,2);not null"`
	ShippingCost    float64        `gorm:"type:decimal(10,2);default:0"`
	RecipientName   string         `gorm:"type:char(36);not null"`
	Phone           string         `gorm:"type:char(36);not null"`
	ShippingAddress string         `gorm:"type:text" json:"shipping_address"`
	Tax             float64        `gorm:"type:decimal(10,2);not null"`
	VoucherCode     *string        `gorm:"type:varchar(100)" json:"voucherCode,omitempty"`
	VoucherDiscount float64        `gorm:"default:0" json:"voucherDiscount"`
	AmountToPay     float64        `gorm:"type:decimal(10,2);not null"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`

	Shipment Shipment    `gorm:"foreignKey:OrderID"`
	Items    []OrderItem `gorm:"foreignKey:OrderID"`
}

type Shipment struct {
	ID           uuid.UUID `gorm:"type:char(36);primaryKey"`
	OrderID      uuid.UUID `gorm:"type:char(36);unique;not null"`
	TrackingCode string    `gorm:"type:varchar(100)"`
	Status       string    `gorm:"type:varchar(20);default:'shipped';check:status IN ('shipped', 'delivered', 'returned')" json:"status"`
	Notes        *string   `gorm:"type:text"`
	ShippedAt    *time.Time
	DeliveredAt  *time.Time
}

type Payment struct {
	ID       uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	UserID   uuid.UUID `gorm:"type:char(36);not null" json:"userId"`
	Fullname string    `gorm:"type:varchar(255);not null" json:"fullname"`
	Email    string    `gorm:"type:varchar(255);not null" json:"email"`
	OrderID  uuid.UUID `gorm:"type:char(36);not null" json:"orderId"`
	Method   string    `gorm:"type:varchar(50);not null" json:"method"`
	Status   string    `gorm:"type:varchar(20);default:'pending';check:status IN ('success', 'pending', 'failed')" json:"status"`
	PaidAt   time.Time `gorm:"autoCreateTime" json:"paidAt"`
	Total    float64   `gorm:"type:decimal(10,2);not null"`

	Order Order `gorm:"foreignKey:OrderID" json:"package"`
	User  User  `gorm:"foreignKey:UserID" json:"user"`
}

type OrderItem struct {
	ID          uuid.UUID      `gorm:"type:char(36);primaryKey"`
	OrderID     uuid.UUID      `gorm:"type:char(36);not null;index"`
	IsReviewed  bool           `gorm:"default:false"`
	ProductID   uuid.UUID      `gorm:"type:char(36);not null"`
	ProductName string         `gorm:"type:varchar(255);not null"`
	ProductSlug string         `gorm:"type:varchar(255);not null"`
	Image       string         `gorm:"type:varchar(255)"`
	Discount    *float64       `gorm:"type:decimal(10,2);default:0" json:"discount"`
	Price       float64        `gorm:"type:decimal(10,2);not null"`
	Quantity    int            `gorm:"not null"`
	Subtotal    float64        `gorm:"type:decimal(10,2)"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Voucher struct {
	ID           uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	Code         string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"code"`
	Description  string         `gorm:"type:text" json:"description"`
	DiscountType string         `gorm:"type:varchar(20);not null" json:"discountType"`
	Discount     float64        `gorm:"not null" json:"discount"`
	MaxDiscount  *float64       `json:"maxDiscount,omitempty"`
	Quota        int            `gorm:"not null" json:"quota"`
	IsReusable   bool           `gorm:"default:false" json:"isReusable"`
	ExpiredAt    time.Time      `gorm:"not null" json:"expiredAt"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

type UsedVoucher struct {
	ID        uuid.UUID      `gorm:"type:char(36);primaryKey"`
	UserID    uuid.UUID      `gorm:"type:char(36);index;not null"`
	VoucherID uuid.UUID      `gorm:"type:char(36);index;not null"`
	UsedAt    time.Time      `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// TRANSACTION SERVICES MODEL ================================

// NOTIFICATIONS SERVICES MODEL ================================

type NotificationType struct {
	ID             uuid.UUID      `gorm:"type:char(36);primaryKey"`
	Code           string         `gorm:"unique;not null"`
	Title          string         `gorm:"type:varchar(255);not null"`
	Category       string         `gorm:"type:varchar(100)"`
	DefaultEnabled bool           `gorm:"default:true"`
	IsRequired     bool           `gorm:"default:false"`
	CreatedAt      time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

type NotificationSetting struct {
	ID                 uuid.UUID      `gorm:"type:char(36);primaryKey"`
	UserID             uuid.UUID      `gorm:"type:char(36);index;not null"`
	NotificationTypeID uuid.UUID      `gorm:"type:char(36);index;not null"`
	Channel            string         `gorm:"type:varchar(20);not null;check:channel IN ('email','browser')"`
	Enabled            bool           `gorm:"default:true"`
	CreatedAt          time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt          time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`

	NotificationType NotificationType `gorm:"foreignKey:NotificationTypeID"`
	User             User             `gorm:"foreignKey:UserID"`
}

type Notification struct {
	ID        uuid.UUID      `gorm:"type:char(36);primaryKey"`
	UserID    uuid.UUID      `gorm:"type:char(36);not null;index"`
	TypeCode  string         `gorm:"type:varchar(100);not null"`
	Title     string         `gorm:"type:varchar(255);not null"`
	Message   string         `gorm:"type:text;not null"`
	Channel   string         `gorm:"type:varchar(50);not null"`
	IsRead    bool           `gorm:"default:false"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// NOTIFICATIONS SERVICES MODEL ================================

func setUUIDIfNil(id *uuid.UUID) {
	if *id == uuid.Nil {
		*id = uuid.New()
	}
}

func (u *User) BeforeCreate(tx *gorm.DB) error                 { setUUIDIfNil(&u.ID); return nil }
func (c *Cart) BeforeCreate(tx *gorm.DB) error                 { setUUIDIfNil(&c.ID); return nil }
func (t *Token) BeforeCreate(tx *gorm.DB) error                { setUUIDIfNil(&t.ID); return nil }
func (o *Order) BeforeCreate(tx *gorm.DB) error                { setUUIDIfNil(&o.ID); return nil }
func (r *Review) BeforeCreate(tx *gorm.DB) error               { setUUIDIfNil(&r.ID); return nil }
func (b *Banner) BeforeCreate(tx *gorm.DB) error               { setUUIDIfNil(&b.ID); return nil }
func (p *Profile) BeforeCreate(tx *gorm.DB) error              { setUUIDIfNil(&p.ID); return nil }
func (p *Product) BeforeCreate(tx *gorm.DB) error              { setUUIDIfNil(&p.ID); return nil }
func (a *Address) BeforeCreate(tx *gorm.DB) error              { setUUIDIfNil(&a.ID); return nil }
func (v *Voucher) BeforeCreate(tx *gorm.DB) error              { setUUIDIfNil(&v.ID); return nil }
func (p *Payment) BeforeCreate(tx *gorm.DB) error              { setUUIDIfNil(&p.ID); return nil }
func (c *Category) BeforeCreate(tx *gorm.DB) error             { setUUIDIfNil(&c.ID); return nil }
func (s *Shipment) BeforeCreate(tx *gorm.DB) error             { setUUIDIfNil(&s.ID); return nil }
func (oi *OrderItem) BeforeCreate(tx *gorm.DB) error           { setUUIDIfNil(&oi.ID); return nil }
func (n *Notification) BeforeCreate(tx *gorm.DB) error         { setUUIDIfNil(&n.ID); return nil }
func (uv *UsedVoucher) BeforeCreate(tx *gorm.DB) error         { setUUIDIfNil(&uv.ID); return nil }
func (g *ProductGallery) BeforeCreate(tx *gorm.DB) error       { setUUIDIfNil(&g.ID); return nil }
func (nt *NotificationType) BeforeCreate(tx *gorm.DB) error    { setUUIDIfNil(&nt.ID); return nil }
func (ns *NotificationSetting) BeforeCreate(tx *gorm.DB) error { setUUIDIfNil(&ns.ID); return nil }
