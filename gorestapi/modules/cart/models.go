package cart

import (
	"gorestapi/db"
	"gorestapi/modules/product"
	"gorestapi/modules/user"
)

type Cart struct {
	db.Model
	User     user.User         `json:"-"`
	UserId   string            `json:"user_id"`
	Products []product.Product `json:"products" gorm:"foreignKey:CartId"`
}

func (c *Cart) Total() float64 {
	var total float64
	for _, item := range c.Products {
		total += item.Price
	}
	return total
}
