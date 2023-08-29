package model

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	Id uint `gorm:"primaryKey;autoIncrement"`
	// foregin key
	Name      string `gorm:"not null; type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Campaign  []Campaign     `gorm:"foreignKey:id"`
}
