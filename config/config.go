package config

import (
	"complaint-service/internal/model"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// user := "admin"
	// password := "root"
	// host := "127.0.0.1"
	// port := "3306"
	// dbname := "customer_db"

	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbname := os.Getenv("MYSQL_DATABASE")

	fmt.Println("MYSQL_USER =", user)
	fmt.Println("MYSQL_PASSWORD =", password)
	fmt.Println("MYSQL_HOST =", host)
	fmt.Println("MYSQL_PORT =", port)
	fmt.Println("MYSQL_DATABASE =", dbname)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)

	fmt.Println("DSN:", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to MySQL:", err)
	}

	if err := db.AutoMigrate(
		&model.Customer{},
		&model.Auth{},
	); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DB = db
	log.Println("Connected to MySQL successfully!")
}
