package api

import (
	v1 "go-gin-learning/internal/api/v1"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/auth", Auth)

	v1.InitRouter(r)
}
