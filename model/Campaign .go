package model

import (
	"time"

	"gorm.io/gorm"
)

type Campaign struct {
	Id uint `gorm:"primaryKey;autoIncrement"`
	// foregin key
	Name        string `gorm:"not null; type:varchar(255)"`
	Description string `gorm:"not null; type:varchar(255)"`
	Target      uint   `gorm:"not null; type:varchar(255)"`
	Start_date  time.Time
	End_date    time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Donation    []Donation     `gorm:"foreignKey:Id_campaign"`
}
