package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rootlulu/go-gin-biu-biu-biu/internal/middlewares/jwt"
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
