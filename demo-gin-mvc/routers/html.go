package routers

import (
	"demo-gin-mvc/controllers/html"

	"github.com/gin-gonic/gin"
)

func RoutersHtml(r *gin.Engine) {
	r.GET("/", html.Html)
}
