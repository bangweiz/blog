package middleware

import (
	"github.com/bangweiz/blog/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int

		code = http.StatusOK
		token := c.Request.Header.Get("Authrization")
		if token == "" {
			code = http.StatusBadRequest
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = http.StatusBadRequest
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = 404
			}
		}

		if code != http.StatusOK {
			data := make(map[string]string, 1)
			data["message"] = "Unauthorized action"
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}