package tokenController

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DefaultController struct{}

func (con DefaultController) Set(c *gin.Context) {
	// 设置 cookie
	c.SetCookie("username", "里斯", 3600, "/", "localhost/", false, false)

	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func (con DefaultController) Get(c *gin.Context) {
	// 获取 cookie
	username, _ := c.Cookie("username")

	c.String(http.StatusOK, "cookie: "+username)
}
