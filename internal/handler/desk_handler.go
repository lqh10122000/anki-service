package handler

import (
	"complaint-service/internal/service"
	response "complaint-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type deskHandler struct {
	service service.DeskService
}

func NewDeskHandler(service service.DeskService) *deskHandler {
	return &deskHandler{service: service}
}

// @Summary add notes
// @Description add notes to the system
// @Tags notes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param notes body model.Notes true "Note data"
// @Success 200 {object} response.ApiResponse
// @Failure 500 {object} response.ApiResponse
// @Router /api/v1/addNotes [post]
func (h *deskHandler) GetAllDesk(c *gin.Context) {
	desks, err := h.service.GetAllDesk()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error("Failed to fetch customers"))
		return
	}

	c.JSON(http.StatusOK, response.Success(desks))
}
