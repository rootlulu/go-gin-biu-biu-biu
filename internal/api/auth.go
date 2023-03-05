package api

import (
	"fmt"
	"log"
	"net/http"

	"go-gin-learning/internal/constant/e"
	"go-gin-learning/internal/pkg/app"
	"go-gin-learning/pkg/util"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// Auth func.
func Auth(c *gin.Context) {
	valid := validation.Validation{}

	var auth struct {
		Username string `valid:"Required; MaxSize(50)"`
		Password string `valid:"Required; MaxSize(50)"`
	}
	c.ShouldBind(&auth)

	ok, _ := valid.Valid(&auth)
	// todo
	if !ok {
		fmt.Println("error")
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
	return true
}
