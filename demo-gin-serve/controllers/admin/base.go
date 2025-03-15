package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseControllers struct{}

func (con BaseControllers) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
