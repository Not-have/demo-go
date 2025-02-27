package main

import (
	routers "demo-gin/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("./template/**")

	// 路由分组
	routers.Base(r)
	routers.Other(r)

	r.Run(":8010") // listen and serve on 0.0.0.0:8080
}
