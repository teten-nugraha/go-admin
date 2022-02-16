package model

import "gorm.io/gorm"

type Order struct {
	Id         uint        `json:"id"`
	FirstName  string      `json:"first_name"`
	LastName   string      `json:"last_name"`
	Email      string      `json:"email"`
	UpdatedAt  string      `json:"updated_at"`
	CreatedAt  string      `json:"created_at"`
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderId"`
}

type OrderItem struct {
	Id           uint    `json:"id"`
	OrderId      uint    `json:"order_id"`
	ProductTitle string  `json:"product_title"`
	Price        float32 `json:"price"`
	Quantity     uint    `json:"quantity"`
}

func (product *Order) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Order{}).Count(&total)

	return total
}

func (product *Order) Take(db *gorm.DB, limit int, offset int) interface{} {
	var orders []Order

	db.Preload("OrderItems").Offset(offset).Limit(limit).Find(&orders)

	return orders
}
