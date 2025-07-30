package model

import "time"

type Customer struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	FirstName string
	LastName  string
	Phone     string
	Email     string
	Address   string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
