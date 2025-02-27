package routers

import (
	"fmt"
	"net/http"

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
	otherRouter.GET("/string", func(c *gin.Context) {
		c.String(http.StatusOK, "值： %v", "我是 字符串返回")
	})

	otherRouter.GET("/object01", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"name":    "我是 对象返回",
			"age":     18,
			"success": true,
			"list": []string{
				"1",
				"2",
				"3",
			},
		})
	})

	otherRouter.GET("/object02", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":    "我是 对象返回",
			"age":     18,
			"success": true,
		})
	})

	otherRouter.GET("/object03", func(c *gin.Context) {
		obj := &User{
			Name: "我是 User 对象返回",
			Age:  18,
		}
		c.JSON(http.StatusOK, obj)
	})

	// http://localhost:8010/jsonp?callback=xx
	otherRouter.GET("/jsonp", func(c *gin.Context) {
		obj := &User{
			Name: "我是 User 对象返回",
			Age:  18,
		}
		c.JSONP(http.StatusOK, obj)
	})

	otherRouter.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"name": "是我 xml",
		})
	})

	// 这个并不常用
	otherRouter.GET("/html", func(c *gin.Context) {
		// 需要配置 otherRouter.LoadHTMLGlob("./template/**")
		c.HTML(http.StatusOK, "index-01.html", gin.H{
			"title": "我是后台的数据",
		})
	})

	// get 传旨
	// http://localhost:8010/get-params?page=10&pageNum=50
	otherRouter.GET("/get-params", func(c *gin.Context) {
		pageNum := c.Query("pageNum")
		page := c.DefaultQuery("page", "1")

		c.JSON(http.StatusOK, gin.H{
			"page":    page,
			"pageNum": pageNum,
		})
	})

	// post 传值
	otherRouter.POST("/post-params", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		// type := c.DefaultPostForm("type", "pc")

		if username != "" && password != "" {
			c.JSON(http.StatusOK, gin.H{
				"userName": username,
				"success":  true,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "Missing username or password",
			})
		}
	})

	// 绑定数据到结构体
	otherRouter.POST("/get-struct", func(c *gin.Context) {
		useInfo := &UserInfo{}

		if err := c.ShouldBind(&useInfo); err == nil {
			fmt.Printf("%#v", useInfo)

			c.JSON(http.StatusOK, useInfo)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "Missing username or password",
			})
		}
	})
}
