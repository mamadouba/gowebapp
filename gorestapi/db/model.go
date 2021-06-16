package db

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Model struct {
	Id        string    `json:"id" gorm:"primary_key, unique"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m Model) SetId() {
	m.Id = uuid.NewV4().String()
}

func GetUid() string {
	return uuid.NewV4().String()
}
