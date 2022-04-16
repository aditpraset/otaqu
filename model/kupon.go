package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type Kupon struct {
	UUID        uuid.UUID       `json:"uuid" gorm:"type:char(36);primary_key"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Rating      decimal.Decimal `json:"rating"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	KuponPrice  []KuponPrice    `gorm:"Foreignkey:KuponID;association_foreignkey:uuid;references:UUID"`
}

type KuponPrice struct {
	UUID      uuid.UUID `json:"uuid" gorm:"type:char(36);primaryKey"`
	KuponID   uuid.UUID `json:"kupon_id" gorm:"type:char(36)"`
	Name      string    `json:"name"`
	Price     uint64    `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Kupon) TableName() string {
	return "kupon"
}

func (KuponPrice) TableName() string {
	return "kupon_price"
}
