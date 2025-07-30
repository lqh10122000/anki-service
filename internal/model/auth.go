package model

import "time"

type Auth struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	Username  string
	Password  string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
