package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rchenhyy/demo-ginex/models"
	"net/http"
)

func GetUser(c *gin.Context) {
	name := c.Param("name")
	age := c.DefaultQuery("age", "unknown")

	c.String(http.StatusOK, fmt.Sprintf("name: %v, age: %v", name, age))
}

func RegisterUser(c *gin.Context) {
	/*
		email := c.PostForm("email")
		password := c.PostForm("password")
		_ = c.DefaultPostForm("password-again", "")
	*/
	var user models.UserModel
	if err := c.ShouldBind(&user); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	c.String(http.StatusOK, fmt.Sprintf("email: %v, password: %v", user.Email, user.Password))
}
