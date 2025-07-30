package model

import "time"

type Desks struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	Name      string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
