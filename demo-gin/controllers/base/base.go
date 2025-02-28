package base

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 结构体
func Home(c *gin.Context) {
	c.String(http.StatusOK, "首页")
}

func List(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "我是 get 请求",
	})
}

func Add(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "我是 add 请求",
	})
}

func Edit(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "我是 put 请求",
	})
}

func Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "我是 delete 请求",
	})
}
