package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Base(r *gin.Engine) {
	baseRouter := r.Group("/base")

	{

		baseRouter.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "首页")
		})

		baseRouter.GET("/list", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "我是 get 请求",
			})
		})

		baseRouter.POST("/add", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "我是 post 请求",
			})
		})

		baseRouter.PUT("/edit", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "我是 put 请求",
			})
		})

		baseRouter.DELETE("/delete", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "我是 delete 请求",
			})
		})
	}
}
