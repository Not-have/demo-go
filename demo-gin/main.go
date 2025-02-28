package main

import (
	routers "demo-gin/routers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func initMiddlewareOne(c *gin.Context) {
	fmt.Println("我是全局中间件")

	// 终止请求
	// c.Abort()

	c.Next()

}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("./template/**")

	r.Use(initMiddlewareOne)

	// 路由分组
	routers.Base(r)
	routers.Other(r)

	r.Run(":8010") // listen and serve on 0.0.0.0:8080
}
