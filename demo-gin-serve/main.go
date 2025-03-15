package main

import (
	"demo-gin-serve/routers"

	"github.com/gin-gonic/gin"
)

func main() {

	// 使用 Default 默认，代表使用了 Logger 和 Recovery 中间件，不想使用的话可以使用 gin.New()
	r := gin.Default()

	r.LoadHTMLGlob("./templates/*")

	r.Static("/static", "./static") // 静态文件目录映射

	// r.Use() 这个里面写的是全局中间件

	routers.RoutersAdminLogin(r)
	routers.RoutersAdminBase(r)

	r.Run()
}
