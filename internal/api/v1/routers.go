package v1

import (
	"go-gin-learning/internal/middlewares/jwt"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	apiv1 := r.Group("/api/v1")

	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/ping", pong)
	}

}

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
