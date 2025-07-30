package router

import (
	"complaint-service/internal/batch"
	"complaint-service/internal/handler"
	"complaint-service/internal/repository"
	"complaint-service/internal/service"
	"complaint-service/middleware"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	customerRepo := repository.NewCustomerRepository()
	customerService := service.NewCustomerService(customerRepo)
	customerHandler := handler.NewCustomerHandler(customerService)

	deskRepo := repository.NewDeskRepository()
	deskService := service.NewDeskService(deskRepo)
	deskHandler := handler.NewDeskHandler(deskService)

	noteRepo := repository.NewNoteRepository()
	noteService := service.NewNoteService(noteRepo, deskRepo, deskService)
	noteHandler := handler.NewNoteHandler(noteService)

	authRepo := repository.NewAuthRepository()
	mailHost := os.Getenv("MAIL_HOST")
	mailPort := os.Getenv("MAIL_PORT")
	mailUser := os.Getenv("MAIL_USER")
	mailPass := os.Getenv("MAIL_PASS")
	mailHostParsed, _ := strconv.Atoi(mailPort)
	mailService := service.NewMailService(mailHost, mailHostParsed, mailUser, mailPass, mailUser)
	authService := service.NewAuthService(authRepo, mailService)
	authHandler := handler.NewAuthHandler(authService)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api/v1")

	api.POST("/addNotes", noteHandler.AddNotes)

	api.GET("/desks", deskHandler.GetAllDesk)

	// Initialize batch processor
	processor := &batch.Processor{Repo: customerRepo}
	batch.Schedule(processor)

	// authentication and authorization routes can be added here
	api.Use(middleware.RateLimitRedisMiddleware(1000, time.Hour))

	api.POST("/login", func(c *gin.Context) {
		if token, err := authHandler.Login(c); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		} else {
			c.JSON(200, gin.H{"token": token})
		}
	})

	api.POST("/register", func(c *gin.Context) {
		if err := authHandler.Register(c); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		}
	})

	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/customers", customerHandler.GetAll)
		api.POST("/customers", customerHandler.Create)
		api.DELETE("/customers/:id", customerHandler.Delete)
		api.PUT("/customers/:id", customerHandler.Update)

	}

	return r
}
