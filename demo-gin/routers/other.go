package routers

import (
	"demo-gin/controllers/other"

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

func Other(r *gin.Engine) {
	otherRouter := r.Group("/other")
	otherRouter.GET("/string", other.String)

	otherRouter.GET("/object01", other.Object01)

	otherRouter.GET("/object02", other.Object02)

	otherRouter.GET("/object03", other.Object03)

	// http://localhost:8010/jsonp?callback=xx
	otherRouter.GET("/jsonp", other.Jsonp)

	otherRouter.GET("/xml", other.Xml)

	// 这个并不常用
	otherRouter.GET("/html", other.Html)

	// get 传旨
	// http://localhost:8010/get-params?page=10&pageNum=50
	otherRouter.GET("/get-params", other.GetParams)

	// post 传值
	otherRouter.POST("/post-params", other.PostParams)

	// 绑定数据到结构体
	otherRouter.POST("/get-struct", other.GetStruct)
}
