package handler

import (
	"complaint-service/internal/model"
	"complaint-service/internal/service"
	response "complaint-service/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	service service.CustomerService
}

func NewCustomerHandler(s service.CustomerService) *CustomerHandler {
	return &CustomerHandler{s}
}

// @Summary Get all customers
// @Description Get all customers with pagination
// @Tags customers
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number"
// @Param size query int false "Page size"
// @Success 200 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Router /api/v1/customers [get]
func (h *CustomerHandler) GetAll(c *gin.Context) {

	// 1. Parse query params
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	sortParam := c.DefaultQuery("sort", "created_at:desc")

	// 2. Parse sort param
	sortField := "created_at"
	sortOrder := "desc"
	if strings.Contains(sortParam, ":") {
		parts := strings.Split(sortParam, ":")
		if len(parts) == 2 {
			sortField = parts[0]
			sortOrder = parts[1]
		}
	}
	orderStr := fmt.Sprintf("%s %s", sortField, sortOrder)
	offset := (page - 1) * size

	// 3. Gọi service: trả về data, tổng số lượng, error
	customers, total, err := h.service.GetAllPaginated(offset, size, orderStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error("Failed to fetch customers"))
		return
	}

	// 4. Trả response chuẩn có pagination
	c.JSON(http.StatusOK, response.PaginatedResponse(customers, total, page, size, sortParam))
}

// @Summary create customer
// @Description Get all customers with pagination
// @Tags customers
// @Accept json
// @Produce json
// @Param customer body model.Customer true "Customer data"
// @Success 200 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Router /api/v1/customers [post]
func (h *CustomerHandler) Create(c *gin.Context) {
	var customer model.Customer
	log.Println("Creating a new customer: ", c.Request.Body)
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, response.Error("Invalid input: "+err.Error()))
		return
	}

	err := h.service.Create(&customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error("Failed to create customer"))
		return
	}

	c.JSON(http.StatusCreated, response.Success(customer))
}

// @Summary create customer
// @Description Get all customers with pagination
// @Tags customers
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param page query int false "customer ID"
// @Success 200 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Router /api/v1/customers [delete]
func (h *CustomerHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	//partse the id to uint
	// parse the id to uint
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error("Invalid customer ID"))
		return
	}

	err = h.service.Delete(uint(uintID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error("Failed to delete customer"))
		return
	}

	c.JSON(http.StatusOK, response.Success("Customer deleted successfully"))
}

// @Summary Update a customer
// @Description Update customer by ID
// @Tags customers
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Customer ID"
// @Param customer body model.Customer true "Customer object"
// @Success 200 {object} response.ApiResponse
// @Failure 400 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Router /api/v1/customers/{id} [put]
func (h *CustomerHandler) Update(c *gin.Context) {
	id := c.Param("id")
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error("Invalid customer ID"))
		return
	}

	var customer model.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, response.Error("Invalid input: "+err.Error()))
		return
	}

	customer.ID = uint(uintID)

	err = h.service.Update(uint(uintID), &customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error("Failed to update customer"))
		return
	}

	c.JSON(http.StatusOK, response.Success(customer))
}
