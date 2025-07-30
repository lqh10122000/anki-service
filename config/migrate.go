package config

import (
	"complaint-service/internal/model"
	"log"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(&model.Notes{}, &model.Desks{})
	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}
	log.Println("✅ Database migrated successfully")
}
