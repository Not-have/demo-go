package routers

import (
	"demo-gin/controllers/other"
	"fmt"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserInfo struct {
	Password string `form:"password" json:"password"`
	UserName string `form:"username" json:"username"`
}

func initMiddlewareOtherOne(c *gin.Context) {
	fmt.Println(c.Request.URL)
	fmt.Println("我是分组中间件 1")

	// 终止请求
	// c.Abort()

	c.Next()

}

func initMiddlewarOtherTwo(c *gin.Context) {
	fmt.Println("我是分组中间件 2")

	// 终止请求
	// c.Abort()

	c.Next()

}

func Other(r *gin.Engine) {
	otherRouter := r.Group("/other", initMiddlewareOtherOne)

	otherRouter.Use(initMiddlewarOtherTwo)

	otherRouter.GET("/string", other.String)

	otherRouter.GET("/object01", other.Object01)

	otherRouter.GET("/object02", other.Object02)

	otherRouter.GET("/object03", other.Object03)

	// http://localhost:8010/jsonp?callback=xx
	otherRouter.GET("/jsonp", other.Jsonp)

	otherRouter.GET("/xml", other.Xml)

	// 这个并不常用
	otherRouter.GET("/html", other.Html)
	otherRouter.GET("/upload", other.Upload)

	// get 传旨
	// http://localhost:8010/get-params?page=10&pageNum=50
	otherRouter.GET("/get-params", other.GetParams)

	// post 传值
	otherRouter.POST("/post-params", other.PostParams)

	// 绑定数据到结构体
	otherRouter.POST("/get-struct", other.GetStruct)
}
