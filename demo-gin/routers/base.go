package routers

import (
	"demo-gin/controllers/base"

	"github.com/gin-gonic/gin"
)

func Base(r *gin.Engine) {
	baseRouter := r.Group("/base")

	{

		baseRouter.GET("/", base.Home)

		baseRouter.GET("/list", base.List)

		baseRouter.POST("/add", base.Add)

		baseRouter.PUT("/edit", base.Edit)

		baseRouter.DELETE("/delete", base.Delete)
	}
}
