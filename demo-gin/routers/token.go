package routers

import (
	tokenController "demo-gin/controllers/token"

	"github.com/gin-gonic/gin"
)

func Cookie(r *gin.Engine) {
	cookieRouter := r.Group("/token")

	{
		cookieRouter.GET("/cookie", tokenController.DefaultController{}.Set)
		cookieRouter.GET("/cookie-get", tokenController.DefaultController{}.Get)
	}
}
