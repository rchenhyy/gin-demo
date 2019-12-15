package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		host := c.Request.Host
		url := c.Request.URL
		method := c.Request.Method

		c.Next()
		log.Printf("%s\t%s\t%s\t%s\t%d\n", time.Now().Format(time.RFC3339), host, method, url, c.Writer.Status())
	}
}

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("username")
		if err != nil {
			c.Abort()
			c.String(http.StatusUnauthorized, "please login first")
			return
		}

		log.Printf("Cookie: %v", cookie)
		c.Next()
	}
}
