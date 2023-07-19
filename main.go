package main

import (
	"genki/controllers"
	"genki/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	r.Static("vendor", "./static/vendor")

	r.LoadHTMLGlob("templates/**/**")

	models.ConnectDatabase()
	models.DBMigrate()

	r.GET("/notes", controllers.NotesIndex)
	r.GET("/notes/new", controllers.NotesNew)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/index.html", gin.H{
			"title": "My awesome Website",
		})
	})

	log.Println("Server startd!")
	r.Run(":8080") // Default Port 8080
}
