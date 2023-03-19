package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rootlulu/go-gin-biu-biu-biu/internal/constant/e"
	_ "github.com/rootlulu/go-gin-biu-biu-biu/internal/middlewares/jwt"
	model "github.com/rootlulu/go-gin-biu-biu-biu/internal/models"
	"github.com/rootlulu/go-gin-biu-biu-biu/internal/pkg/app"
	"github.com/rootlulu/go-gin-biu-biu-biu/pkg/logging"
)

func InitRouter(r *gin.Engine) {
	apiv1 := r.Group("/api/v1")

	// apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/ping", pong)
		apiv1.GET("/users", users)
		apiv1.GET("/userName/:name", usersWithName)
		apiv1.GET("/userId", usersWithId)
		apiv1.GET("/userPassword", usersWithPassword)
	}
}

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, app.Response{
		e.SUCCESS,
		e.CodeMsg(e.SUCCESS),
		map[string]string{"message": "pong"},
	})
}

func users(c *gin.Context) {
	// curl -XGET http://127.0.0.1:8000/api/v1/users -d '{"token": "xxx"}'
	rows, err := model.Query("SELECT * FROM lulu;")
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

func usersWithName(c *gin.Context) {
	// queryUri.
	// curl -XGET http://127.0.0.1:8000/api/v1/userName/lulu -d '{"token": "xxx"}'
	var nameS struct {
		Name string `uri:"name" binding:"required"`
	}
	var id int
	var name, password string
	c.ShouldBindUri(&nameS)
	err := model.QueryRow(
		fmt.Sprintf("SELECT * FROM lulu where name='%s'", nameS.Name),
	).Scan(&id, &name, &password)
	if err != nil {
		logging.Error(err)
		c.JSON(http.StatusOK, app.Response{e.ERROR, e.CodeMsg(e.ERROR), err})
		return
	}
	c.JSON(http.StatusOK, app.Response{e.SUCCESS, e.CodeMsg(e.SUCCESS),
		model.User{name, password}})
}

func usersWithId(c *gin.Context) {
	// queryString
	// curl -XGET http://127.0.0.1:8000/api/v1/userId?id=2
	var idS struct {
		Id int `form:"id"`
	}
	var id int
	var name, password string
	c.ShouldBindQuery(&idS)
	err := model.QueryRow(
		fmt.Sprintf("SELECT * FROM lulu where id=%d", idS.Id),
	).Scan(&id, &name, &password)
	if err != nil {
		logging.Error(err)
		c.JSON(http.StatusOK, app.Response{e.ERROR, e.CodeMsg(e.ERROR), err})
		return
	}
	c.JSON(http.StatusOK, app.Response{e.SUCCESS, e.CodeMsg(e.SUCCESS),
		model.User{name, password}})
}

func usersWithPassword(c *gin.Context) {
	// queryJson
	// curl -XGET http://127.0.0.1:8000/api/v1/userPassword -d '{"password": "Judi"}'
	var passwordS struct {
		Password string `json:"password"`
	}
	var id int
	var name, password string
	c.ShouldBindJSON(&passwordS)
	err := model.QueryRow(fmt.Sprintf("SELECT * FROM lulu where password='%s'", passwordS.Password)).Scan(&id, &name, &password)
	if err != nil {
		logging.Error(err)
		c.JSON(http.StatusOK, app.Response{e.ERROR, e.CodeMsg(e.ERROR), err})
		return
	}
	c.JSON(http.StatusOK, app.Response{e.SUCCESS, e.CodeMsg(e.SUCCESS), model.User{name, password}})
}
