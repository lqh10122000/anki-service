package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowForm(c *gin.Context) {
	c.HTML(http.StatusOK, "add_note.html", nil)
}

func HandleForm(c *gin.Context) {
	word := c.PostForm("word")
	translate := c.PostForm("translateWord")
	lang := c.PostForm("lang")

	// Tạm thời in ra để test
	c.String(http.StatusOK, fmt.Sprintf("Word: %s, Translate: %s, Lang: %s", word, translate, lang))
}
