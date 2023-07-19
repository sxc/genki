package main

import (
	"genki/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectDatabase() {
	database, err := gorm.Open(mysql.Open("fishmans:fishmans@tcp(127.0.0.1:3306)/note?charset=utf8"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	DB = database
}

func dbMigrate() {
	DB.AutoMigrate(&models.Note{})
}

func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	r.Static("vendor", "./static/vendor")

	r.LoadHTMLGlob("templates/**/**")

	connectDatabase()
	dbMigrate()

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/index.html", gin.H{
			"title": "My awesome Website",
		})
	})

	log.Println("Server startd!")
	r.Run(":8080") // Default Port 8080
}
