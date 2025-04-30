package routers

import (
	"demo-gin-mvc-verification-code/controllers/admin"
	"demo-gin-mvc-verification-code/middlewares"

	"github.com/gin-gonic/gin"
)

func RoutersAdminLogin(r *gin.Engine) {
	// r.Group 第二个参数的意思是中间键，用作路由控制器，等同于 r.Use(middlewares.InitMiddleware)
	routerAdmin := r.Group("/admin", middlewares.InitMiddleware)

	{
		routerAdmin.POST("/login", admin.LoginControllers{}.Index)
		routerAdmin.GET("/user", admin.UserControllers{}.Index)
	}
}
