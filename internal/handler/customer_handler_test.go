package handler_test

import (
	"complaint-service/internal/handler"
	"complaint-service/internal/model"
	"complaint-service/internal/service/mocks"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCustomerHandler_GetAll(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(mocks.CustomerService)
	mockCustomers := []model.Customer{
		{ID: 1, FirstName: "John", LastName: "Doe"},
		{ID: 2, FirstName: "Jane", LastName: "Smith"},
	}
	mockService.On("GetAllPaginated", 0, 10, "created_at desc").Return(mockCustomers, int64(2), nil)

	h := handler.NewCustomerHandler(mockService)

	// Tạo router và gắn handler
	r := gin.Default()
	r.GET("/api/v1/customers", h.GetAll)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/customers?page=1&size=10&sort=created_at:desc", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var result struct {
		Message string `json:"message"`
		Data    any    `json:"data"`
	}
	json.NewDecoder(resp.Body).Decode(&result)
	assert.Equal(t, "Success", result.Message)
	// dataMap, ok := result.Data.(map[string]interface{})
	// if !ok {
	// 	t.Fatalf("expected data to be a map[string]interface{}, got %T", result.Data)
	// }
	// assert.Len(t, dataMap, 2)
}
