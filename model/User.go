package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id               uint   `gorm:"primaryKey;autoIncrement"`
	Email            string `gorm:"not null; unique; type:varchar(255)"`
	Password         string `gorm:"not null; type:varchar(255)"`
	Fullname         string `gorm:"not null; type:varchar(255)"`
	Image            string `gorm:"not null; type:varchar(255)"`
	Handphone_number string `gorm:"not null; type:varchar(255)"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	Donation         []Donation     `gorm:"foreignKey:id_users"`
}
