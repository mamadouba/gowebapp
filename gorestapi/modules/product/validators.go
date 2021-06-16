package product

import "time"

type createProductFields struct {
	Title       string    `validate:"required;regexp=^[a-zA-Z0-9]{2,30}$"`
	Description string    `validate:"required"`
	Price       float64   `validate:"required"`
	Stock       int64     `validate:"omitempty"`
	Discount    float64   `validate:"omitempty"`
	DiscountEnd time.Time `validate:"omitempty"`
	Tags        string    `validate:"omitempty"`
	ImageUrls   string    `validate:"omitempty"`
	CategoryId  string    `json:"category_id" validate:"required"`
}

type updateProductFields struct {
	Title       string    `validate:"omitempty;regexp=^[a-zA-Z0-9]{2,30}$"`
	Description string    `validate:"omitempty"`
	Price       float64   `validate:"omitempty"`
	Stock       int64     `validate:"omitempty"`
	Discount    float64   `validate:"omitempty"`
	DiscountEnd time.Time `validate:"omitempty"`
	Tags        string    `validate:"omitempty"`
	ImageUrls   string    `validate:"omitempty"`
}

func createProductValidator() interface{} {
	return &createProductFields{}
}

func updateProductValidator() interface{} {
	return &updateProductFields{}
}
