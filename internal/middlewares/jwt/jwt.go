package jwt

import (
	"go-gin-learning/internal/constant/e"
	"go-gin-learning/pkg/util"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWT is the middleware to check the token is valid or not.
func JWT() gin.HandlerFunc {
	f := func(c *gin.Context) {
		var data interface{}
		code := e.SUCCESS
		var token struct {
			Token string `json:"token"`
		}
		c.ShouldBindJSON(&token)
		if token.Token == "" {
			code = e.INVALID_PARAMS
		} else {
			_, err := util.ParseToken(token.Token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.AUTH_TOKEN_TIMEOUT
				default:
					code = e.TOKEN_ERROR
				}
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.CodeMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}

	return f
}
