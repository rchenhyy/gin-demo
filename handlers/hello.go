package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloGin(c *gin.Context) {
	c.String(http.StatusOK, "Hello gin!")
}
