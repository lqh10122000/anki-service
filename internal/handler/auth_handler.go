package handler

import (
	"complaint-service/internal/event"
	"complaint-service/internal/service"
	response "complaint-service/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type authHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) *authHandler {
	return &authHandler{service: service}
}

// @Summary login
// @Description Get all customers with pagination
// @Tags auth
// @Param request body LoginRequest true "Register Request Body"
// @Success 200 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Router /api/v1/login [post]
func (h *authHandler) Login(c *gin.Context) (string, error) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Error binding JSON: %v", err)
		return "", err
	}
	log.Print("Login request received: " + req.Username + " password: " + req.Password)

	serviceReq := service.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	}
	token, err := h.service.Login(serviceReq)

	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return "", err
	}
	if token == "" {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return "", err
	}

	event.LogLoginEvent(req.Username)

	c.JSON(200, gin.H{"token": token})

	return "", err
}

// @Summary register customer
// @Description Get all customers with pagination
// @Tags auth
// @Param request body LoginRequest true "Register Request Body"
// @Success 200 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Router /api/v1/register [post]
func (h *authHandler) Register(c *gin.Context) error {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Error binding JSON: %v", err)
		return err
	}

	if req.Username == "" || req.Password == "" {
		c.JSON(http.StatusInternalServerError, response.Error("Failed to fetch customers"))
	}

	serviceReq := service.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	}
	error := h.service.Register(serviceReq)
	if error != nil {
		log.Printf("Error registering user: %v", error)
		c.JSON(http.StatusInternalServerError, response.Error("Failed to register user"))
		return error
	}

	// log.Printf("Register request received: %s with hashed password: %s", req.Username, string(hashedPassword))
	return nil
}
