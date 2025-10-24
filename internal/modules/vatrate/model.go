package vatrate

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VatRate struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Year      int            `json:"year" gorm:"not null"`
	Month     int            `json:"month" gorm:"not null"`
	Rate      float64        `json:"rate" gorm:"not null"`
	IsActive  bool           `json:"is_active" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at" gorm:"not null;default:now()"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"not null;default:now()"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (VatRate) TableName() string {
	return "vat_rates"
}
