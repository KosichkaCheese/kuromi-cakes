package models

import "time"

type DeliveryType struct {
	ID   int    `json:"id" gorm:"primaryKey;autoIncrement;not null;unique"`
	Name string `json:"name" gorm:"not null"`
}

type PaymentType struct {
	ID   int    `json:"id" gorm:"primaryKey;autoIncrement;not null;unique"`
	Name string `json:"name" gorm:"not null"`
}

type Order struct {
	ID             int          `json:"id" gorm:"primaryKey;autoIncrement;not null;unique"`
	DeliveryTypeID int          `gorm:"not null" json:"-"`
	DeliveryType   DeliveryType `json:"delivery_type" gorm:"foreignKey:DeliveryTypeID"`
	Name           string       `json:"name" gorm:"not null"`
	Address        string       `json:"address"`
	Phone          string       `json:"phone" gorm:"not null"`
	PaymentTypeID  int          `gorm:"not null" json:"-"`
	PaymentType    PaymentType  `json:"payment_type" gorm:"foreignKey:PaymentTypeID"`
	Email          string       `json:"email" gorm:"not null"`
	TotalPrice     float64      `json:"price" gorm:"type:decimal(10,2);not null"`
	CreatedAt      time.Time
}

type OrderPost struct {
	DeliveryTypeID int     `json:"delivery_type_id" gorm:"not null"`
	Name           string  `json:"name" gorm:"not null"`
	Address        string  `json:"address" gorm:"not null"`
	Phone          string  `json:"phone" gorm:"not null"`
	PaymentTypeID  int     `json:"payment_type_id" gorm:"not null"`
	Email          string  `json:"email" gorm:"not null"`
	TotalPrice     float64 `json:"price" gorm:"type:decimal(10,2);not null"`
}

type OrderContent struct {
	OrderID  int `json:"order_id" gorm:"primaryKey"`
	CakeID   int `json:"cake_id" gorm:"primaryKey"`
	Quantity int `json:"quantity" gorm:"not null"`

	Order Order `json:"-" gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;"`
	Cake  Cake  `json:"cake" gorm:"foreignKey:CakeID"`
}

type OrderContentPost struct {
	OrderID int `json:"order_id" binding:"required"`
	Items   []struct {
		CakeID   int `json:"cake_id" binding:"required"`
		Quantity int `json:"quantity" binding:"required,min=1"`
	} `json:"items" binding:"required,dive"`
}
