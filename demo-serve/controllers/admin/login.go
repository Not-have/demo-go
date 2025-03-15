package admin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginControllers struct{}

func (con LoginControllers) Index(c *gin.Context) {
	// 使用中间键的数据
	data, _ := c.Get("request")
	fmt.Println(data)
	c.String(http.StatusOK, "login")
}
