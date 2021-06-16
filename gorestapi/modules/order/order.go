package order

import (
	"gorestapi/db"
	"gorestapi/modules/user"
)

type Order struct {
	db.Model
	User   user.User
	UserId string `json:"user_id"`
	Status string `json:"status"`
	Items  []Item `json:"items" gorm:"foreignKey:OrderId"`
}

type Item struct {
	db.Model
	PrductId string `json:"product_id"`
	Quantity int64  `json:"quantity"`
	OrderId  string `json:"order_id"`
}
