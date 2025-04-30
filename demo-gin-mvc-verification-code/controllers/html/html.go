package html

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Html(c *gin.Context) {
	// 需要配置 otherRouter.LoadHTMLGlob("./template/**")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Hello, Gin!",
	})
}
