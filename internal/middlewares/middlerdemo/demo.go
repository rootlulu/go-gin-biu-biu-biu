package middlerdemo

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// V1Others ...
func V1Others() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf := make([]byte, 1024)
		c.Copy().Request.Body.Read(buf)
		fmt.Printf("The Headers: %v \n The body: %v \n", c.Request.Header, string(buf))
		c.Next()
	}
}
