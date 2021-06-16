package category

import (
	"gorestapi/db"
)

type Category struct {
	db.Model
	Name string `json:"name"`
}
