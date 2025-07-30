package model

import "time"

type Notes struct {
	ID            uint `gorm:"primaryKey;autoIncrement"`
	DeskID        uint
	ModelName     string
	Word          string
	Lang          string
	TranslateWord string
	AudioFilename string
	Tags          string
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}
