package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rootlulu/go-gin-biu-biu-biu/internal/api"
	"github.com/rootlulu/go-gin-biu-biu-biu/internal/config"
	model "github.com/rootlulu/go-gin-biu-biu-biu/internal/models"
	"github.com/rootlulu/go-gin-biu-biu-biu/pkg/logging"
	"github.com/rootlulu/go-gin-biu-biu-biu/pkg/util"

	"github.com/gin-gonic/gin"
)

func init() {
	config.Init()
	util.Init()
	logging.Init()
	model.Init()

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
