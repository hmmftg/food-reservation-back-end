package params

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        string `gorm:"primarykey"`
	Name      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type FoodReservationParams struct {
	StaticBaseUrl string
	StaticPath    string
}
