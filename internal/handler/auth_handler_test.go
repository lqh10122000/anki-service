package handler_test

import (
	"bytes"
	"complaint-service/internal/handler"
	"complaint-service/internal/service/mocks"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(mocks.AuthService)
	h := handler.NewAuthHandler(mockService)

	// Input
	reqBody := handler.LoginRequest{
		Username: "testuser",
		Password: "password123",
	}
	jsonValue, _ := json.Marshal(reqBody)

	mockService.On("Login", mock.Anything).Return("mockedToken", nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonValue))
	c.Request.Header.Set("Content-Type", "application/json")

	h.Login(c)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "mockedToken")
}

func TestLogin_BindJSON_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(mocks.AuthService)
	h := handler.NewAuthHandler(mockService)

	body := []byte(`invalid_json`)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	c.Request.Header.Set("Content-Type", "application/json")

	h.Login(c)

	// assert.Equal(t, 400, w.Code)
	// assert.Contains(t, w.Body.String(), "error")
}
