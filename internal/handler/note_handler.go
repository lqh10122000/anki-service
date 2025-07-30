package handler

import (
	"complaint-service/internal/model"
	"complaint-service/internal/service"
	response "complaint-service/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type noteHandler struct {
	service service.NoteService
}

func NewNoteHandler(service service.NoteService) *noteHandler {
	return &noteHandler{service: service}
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
func (h *noteHandler) AddNotes(c *gin.Context) {
	var note model.Notes

	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, response.Error("Invalid input: "+err.Error()))
		return
	}

	log.Println("adding new note: ", note)

	err := h.service.AddNotes(&note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error("Failed to create customer"))
		return
	}

	c.JSON(http.StatusCreated, response.Success(note))
}
