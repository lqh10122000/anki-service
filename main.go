// @title Complaint Service API
// @version 1.0
// @description This is a sample API for managing complaints.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"complaint-service/config"
	_ "complaint-service/docs"
	"complaint-service/router"
	"embed"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//go:embed templates/*
var templatesFS embed.FS

func main() {
	config.ConnectDB()
	// config.ConnectRedis()
	// config.MigrateDB(config.DB)

	// cfg := config.LoadKafkaConfig()
	// event.InitKafkaLogger(cfg)

	// go event.StartKafkaConsumer(cfg)

	r := router.SetupRouter()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	tmpl := template.Must(template.ParseFS(templatesFS, "templates/*"))
	r.SetHTMLTemplate(tmpl)

	// Giao diện thêm note
	// r.GET("/add-note", handler.ShowForm)
	// r.POST("/add-note", handler.HandleForm)

	r.GET("/add-note", func(c *gin.Context) {
		c.HTML(http.StatusOK, "add_note.html", nil)
	})

	log.Fatal(r.Run(":8080"))
}
