package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(c *gin.Context) {
	name := c.Param("name")
	age := c.DefaultQuery("age", "unknown")

	c.String(http.StatusOK, fmt.Sprintf("name: %v, age: %v", name, age))
}

func NewUser(c *gin.Context) {

}