package product

import (
	"gorestapi/db"
	"gorestapi/modules/category"
	"time"
)

type Product struct {
	db.Model
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Price       float64           `json:"price"`
	Stock       int64             `json:"stock"`
	Discount    float64           `json:"discount"`
	DiscountEnd time.Time         `json:"discount_end"`
	ImageUrls   string            `json:"image_urls"`
	CategoryId  string            `json:"category_id"`
	Category    category.Category `json:"-" gorm:"foreignKey:CategoryId"`
	Comments    []Comment         `json:"-" gorm:"foreignKey:ProductId"`
	Ratings     []Rating          `json:"-" gorm:"foreignKey:ProductId"`
}

type Comment struct {
	db.Model
	UserId    string `json:"user_id"`
	ProductId string `json:"product_id"`
	Message   string `json:"message"`
}

type Rating struct {
	db.Model
	UserId    string `json:"user_id"`
	ProductId string `json:"product_id"`
	Value     int64  `json:"rate"`
}
