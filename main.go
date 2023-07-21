package main

import (
	"genki/controllers"
	"genki/middlewares"
	"genki/models"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	r.Static("vendor", "./static/vendor")

	r.LoadHTMLGlob("templates/**/**")

	models.ConnectDatabase()
	models.DBMigrate()

	// Sessions INit
	store := memstore.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("notes", store))

	r.Use(middlewares.AuthenticateUser())

	// Route group - notes
	notes := r.Group("/notes")
	{
		notes.GET("/", controllers.NotesIndex)
		notes.GET("/new", controllers.NotesNew)
		notes.POST("/", controllers.NotesCreate)
		notes.GET("/:id", controllers.NotesShow)
		notes.GET("/edit/:id", controllers.NotesEditPage)
		notes.POST("/:id", controllers.NotesUpdate)
		notes.DELETE("/:id", controllers.NotesDelete)
	}

	r.GET("/login", controllers.LoginPage)
	r.GET("/signup", controllers.SignupPage)

	r.POST("/login", controllers.Login)
	r.POST("/signup", controllers.Signup)
	r.POST("/logout", controllers.Logout)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/index.html", gin.H{
			"title":     "My awesome Website",
			"logged_in": (c.GetUint64("user_id") > 0),
		})
	})

	log.Println("Server startd!")
	r.Run(":3000")
}
