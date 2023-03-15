package api

import (
	"database/sql"
	"net/http"

	"github.com/rootlulu/go-gin-biu-biu-biu/internal/config"
	"github.com/rootlulu/go-gin-biu-biu-biu/internal/constant/e"
	"github.com/rootlulu/go-gin-biu-biu-biu/internal/pkg/app"
	"github.com/rootlulu/go-gin-biu-biu-biu/pkg/logging"
	"github.com/rootlulu/go-gin-biu-biu-biu/pkg/util"

	"github.com/gin-gonic/gin"
)

// Auth func.
func Auth(c *gin.Context) {

	var auth struct {
		Username string `valid:"Required; MaxSize(50)"`
		Password string `valid:"Required; MaxSize(50)"`
	}
	err := c.ShouldBind(&auth)
	if err != nil {
		c.JSON(http.StatusOK, app.Response{e.INVALID_PARAMS, e.CodeMsg(e.INVALID_PARAMS), err})
	}

	token, err := util.GenerateToken(auth.Username, auth.Password)

	if err != nil {
		c.JSON(http.StatusOK, app.Response{e.INVALID_PARAMS, e.CodeMsg(e.INVALID_PARAMS), err})
	} else {
		if check(auth) {
			c.JSON(http.StatusOK, app.Response{
				Code: e.SUCCESS,
				Msg:  e.CodeMsg(e.SUCCESS),
				Data: map[string]string{"token": token},
			})
		}
	}

}

// todo
func check(auth interface{}) bool {
	// todo
	db, err := sql.Open(config.DB.Type, config.DB.Path+config.DB.File)
	if err != nil {
		logging.Error("Cant open the db file.", err)
	}
	rows, err := db.Query("SELECT * from lulu;")
	if err != nil {
		logging.Error(err)
	}
	for rows.Next() {
		var id int
		var username string
		var password string

		err := rows.Scan(&id, &username, &password)
		if err != nil {
			logging.Error(err)
		}
	}
	return true
}
