package api

import (
	v1 "github.com/rootlulu/go-gin-biu-biu-biu/internal/api/v1"
	"github.com/rootlulu/go-gin-biu-biu-biu/internal/middlewares/middlerdemo"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlerdemo.V1Others())

	r.POST("/auth", Auth)

	v1.InitRouter(r)
}
