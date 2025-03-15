package routers

import (
	"demo-gin-serve/controllers/admin"

	"github.com/gin-gonic/gin"
)

func RoutersAdminBase(r *gin.Engine) {
	// r.Group 第二个参数的意思是中间键，用作路由控制器，等同于 r.Use(middlewares.InitMiddleware)
	routerAdmin := r.Group("/base")

	{
		routerAdmin.GET("/", admin.BaseControllers{}.Index)
	}
}
