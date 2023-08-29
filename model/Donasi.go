package model

import (
	"time"

	"gorm.io/gorm"
)

type Donation struct {
	Id uint `gorm:"primaryKey;autoIncrement"`
	// foregin key
	Id_campaign    uint   `gorm:"not null"`
	Id_users       uint   `gorm:"not null"`
	Amount         int    `gorm:"not null"`
	Payment_method string `gorm:"not null; type:varchar(255)"`
	Status         string `gorm:"not null; type:varchar(255)"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
