package controllers

import (
	"genki/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotesIndex(c *gin.Context) {
	notes := models.NotesAll()
	c.HTML(
		http.StatusOK,
		"notes/index.html",
		gin.H{
			"notes": notes,
		},
	)
}

func NotesNew(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"notes/new.html",
		gin.H{},
	)
}
