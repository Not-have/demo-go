package routers

import (
	"demo-gin/controllers/base"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func initMiddlewareOne(c *gin.Context) {
	start := time.Now().UnixNano()
	fmt.Println("记录日志1-1")

	// 终止请求
	// c.Abort()

	c.Next()

	fmt.Println("记录日志1-2")
	end := time.Now().UnixNano()

	fmt.Println(end - start)
}

func initMiddlewareTwo(c *gin.Context) {
	start := time.Now().UnixNano()
	fmt.Println("记录日志2-1")

	// 终止请求
	// c.Abort()

	c.Next()

	fmt.Println("记录日志2-1")
	end := time.Now().UnixNano()

	fmt.Println(end - start)
}

func Base(r *gin.Engine) {
	baseRouter := r.Group("/base")

	{

		baseRouter.GET("/", initMiddlewareOne, initMiddlewareTwo, func(c *gin.Context) {
			fmt.Println("我是首页")

			c.String(http.StatusOK, "首页")
		})

		baseRouter.GET("/list", base.List)

		baseRouter.POST("/add", base.Add)

		baseRouter.PUT("/edit", base.Edit)

		baseRouter.DELETE("/delete", base.Delete)
	}
}
