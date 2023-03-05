package main

import (
	"fmt"
	"go-gin-learning/internal/api"
	"go-gin-learning/internal/config"
	model "go-gin-learning/internal/models"
	"go-gin-learning/pkg/logging"
	"go-gin-learning/pkg/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	logging.Init()
	config.Init()
	model.Init()
	logging.Init()
	util.Init()
}

func main() {
	gin.SetMode(config.App.RunMode)
	r := gin.New()
	api.Init(r)

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.App.RunPort),
		Handler:        r,
		ReadTimeout:    time.Duration(config.App.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(config.App.WriteTimeout) * time.Second,
		MaxHeaderBytes: config.App.MaxHeaderBytes,
	}
	server.ListenAndServe()
}
