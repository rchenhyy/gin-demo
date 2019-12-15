package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rchenhyy/demo-ginex/models"
	"log"
	"net/http"
)

func User(c *gin.Context) {
	name := c.Param("name")
	age := c.DefaultQuery("age", "unknown")

	c.String(http.StatusOK, fmt.Sprintf("name: %v, age: %v", name, age))
}

func UserRegister(c *gin.Context) {
	/*
		email := c.PostForm("email")
		password := c.PostForm("password")
		_ = c.DefaultPostForm("password-again", "")
	*/
	var user models.UserModel
	if err := c.ShouldBind(&user); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	passwordAgain := c.DefaultPostForm("password-again", "")
	if user.Password != passwordAgain {
		c.String(http.StatusBadRequest, "password mismatch")
		return
	}

	user.Save()
	c.String(http.StatusOK, fmt.Sprintf("email: %v, password: %v", user.Email, user.Password))
}

func UserLogin(c *gin.Context) {
	var user models.UserModel
	if err := c.ShouldBind(&user); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	u := user.QueryOne()
	if u.Password == user.Password {
		log.Println("login success")
		// set cookie
		c.SetCookie("username", user.Email, 0, "path", c.Request.Host, false, true)
		c.String(http.StatusOK, "user: email=%v", user.Email)
		return
	}

	c.String(http.StatusUnauthorized, "user: email=%v", user.Email)
	// c.Redirect()
}

func UserAvatarUpload() {
	// TO.DO. ...
}
