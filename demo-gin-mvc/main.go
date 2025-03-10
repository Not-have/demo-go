package main

import (
	"demo-gin-mvc/routers"

	"github.com/gin-gonic/gin"
)

func main() {

	// 使用 Default 默认，代表使用了 Logger 和 Recovery 中间件，不想使用的话可以使用 gin.New()
	r := gin.Default()

	// r.Use() 这个里面写的是全局中间件

	routers.RoutersAdminLogin(r)

	r.Run()

}
