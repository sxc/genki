package controllers

import (
	"fmt"
	"genki/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NotesIndex(c *gin.Context) {
	// currentUser := helpers.GetUserFromRequest(c)
	userID := c.GetUint64("user_id")
	var currentUser *models.User
	if userID > 0 {
		currentUser = models.UserFind(userID)
	} else {
		currentUser = nil
	}
	if currentUser == nil || currentUser.ID == 0 {
		c.HTML(
			http.StatusUnauthorized,
			"notes/index.html",
			gin.H{
				"alert": "Unauthorized Access!",
			},
		)
		return
	}
	notes := models.NotesAll(currentUser)
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

type FormData struct {
	Name    string `form:"name"`
	Content string `form:"content"`
}

func NotesCreate(c *gin.Context) {
	// currentUser := helpers.GetUserFromRequest(c)
	userID := c.GetUint64("user_id")
	var currentUser *models.User
	if userID > 0 {
		currentUser = models.UserFind(userID)
	} else {
		currentUser = nil
	}
	if currentUser == nil || currentUser.ID == 0 {
		c.HTML(
			http.StatusUnauthorized,
			"notes/index.html",
			gin.H{
				"alert": "Unauthorized Access!",
			},
		)
		return
	}
	var data FormData
	c.Bind(&data)
	models.NoteCreate(currentUser, data.Name, data.Content)
	c.Redirect(http.StatusMovedPermanently, "/notes")
}

func NotesShow(c *gin.Context) {
	// currentUser := helpers.GetUserFromRequest(c)
	userID := c.GetUint64("user_id")
	var currentUser *models.User
	if userID > 0 {
		currentUser = models.UserFind(userID)
	} else {
		currentUser = nil
	}
	if currentUser == nil || currentUser.ID == 0 {
		c.HTML(
			http.StatusUnauthorized,
			"notes/index.html",
			gin.H{
				"alert": "Unauthorized Access!",
			},
		)
		return
	}
	idstr := c.Param("id")
	id, err := strconv.ParseUint(idstr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	note := models.NotesFind(currentUser, id)
	c.HTML(
		http.StatusOK,
		"notes/show.html",
		gin.H{
			"note": note,
		},
	)
}

func NotesEditPage(c *gin.Context) {
	// currentUser := helpers.GetUserFromRequest(c)
	userID := c.GetUint64("user_id")
	var currentUser *models.User
	if userID > 0 {
		currentUser = models.UserFind(userID)
	} else {
		currentUser = nil
	}
	if currentUser == nil || currentUser.ID == 0 {
		c.HTML(
			http.StatusUnauthorized,
			"notes/index.html",
			gin.H{
				"alert": "Unauthorized Access!",
			},
		)
		return
	}
	idstr := c.Param("id")
	id, err := strconv.ParseUint(idstr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	note := models.NotesFind(currentUser, id)
	c.HTML(
		http.StatusOK,
		"notes/edit.html",
		gin.H{
			"note": note,
		},
	)
}

func NotesUpdate(c *gin.Context) {
	// currentUser := helpers.GetUserFromRequest(c)
	userID := c.GetUint64("user_id")
	var currentUser *models.User
	if userID > 0 {
		currentUser = models.UserFind(userID)
	} else {
		currentUser = nil
	}
	if currentUser == nil || currentUser.ID == 0 {
		c.HTML(
			http.StatusUnauthorized,
			"notes/index.html",
			gin.H{
				"alert": "Unauthorized Access!",
			},
		)
		return
	}
	idstr := c.Param("id")
	id, err := strconv.ParseUint(idstr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	note := models.NotesFind(currentUser, id)
	name := c.PostForm("name")
	content := c.PostForm("content")
	note.Update(name, content)
	c.Redirect(http.StatusMovedPermanently, "/notes/"+idstr)
}

func NotesDelete(c *gin.Context) {
	// currentUser := helpers.GetUserFromRequest(c)
	userID := c.GetUint64("user_id")
	var currentUser *models.User
	if userID > 0 {
		currentUser = models.UserFind(userID)
	} else {
		currentUser = nil
	}
	if currentUser == nil || currentUser.ID == 0 {
		c.HTML(
			http.StatusUnauthorized,
			"notes/index.html",
			gin.H{
				"alert": "Unauthorized Access!",
			},
		)
		return
	}
	idstr := c.Param("id")
	id, err := strconv.ParseUint(idstr, 10, 64)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	models.NotesMarkDelete(currentUser, id)
	c.Redirect(http.StatusSeeOther, "/notes")
}
