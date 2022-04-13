package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type GetKupon struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Rating      string `json:"rating"`
	Price       string `json:"price"`
}

type Kupon struct {
	UUID        uuid.UUID  `json:"uuid" gorm:"type:char(36);primary_key"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Rating      float32    `json:"rating"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	KuponPrice  KuponPrice `json:"kupon_price" gorm:"references:KuponID"`
}

type KuponPrice struct {
	UUID    uuid.UUID `json:"uuid" gorm:"type:char(36);primaryKey"`
	KuponID uuid.UUID `json:"kupon_id" gorm:"type:char(36)"`
	// Kupon     Kupon     `json:"kupon" gorm:"references:UUID"`
	Price     uint64    `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
