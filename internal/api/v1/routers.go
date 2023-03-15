package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rootlulu/go-gin-biu-biu-biu/internal/config"
	"github.com/rootlulu/go-gin-biu-biu-biu/internal/constant/e"
	"github.com/rootlulu/go-gin-biu-biu-biu/internal/middlewares/jwt"
	model "github.com/rootlulu/go-gin-biu-biu-biu/internal/models"
	"github.com/rootlulu/go-gin-biu-biu-biu/internal/pkg/app"
	"github.com/rootlulu/go-gin-biu-biu-biu/pkg/logging"
	"github.com/rootlulu/go-gin-biu-biu-biu/pkg/util"
)

func InitRouter(r *gin.Engine) {
	apiv1 := r.Group("/api/v1")

	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/ping", pong)
		apiv1.GET("/users", users)
	}

}

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func users(c *gin.Context) {
	db := util.SqliteIns{"sqlite3", config.DB.Path + config.DB.File}
	rows, err := db.Query("SELECT * FROM lulu;")
	if err != nil {
		logging.Error(err)
	}
	var result = make([]model.User, 0)
	for rows.Next() {
		var id int
		var username, password string
		rows.Scan(&id, &username, &password)
		result = append(result, model.User{username, password})
	}
	c.JSON(http.StatusOK, app.Response{e.SUCCESS, e.CodeMsg(e.SUCCESS), result})
}
