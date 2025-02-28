package other

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

func String(c *gin.Context) {
	c.String(http.StatusOK, "值： %v", "我是 字符串返回")
}

func Object01(c *gin.Context) {
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
}

func Object02(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":    "我是 对象返回",
		"age":     18,
		"success": true,
	})
}

func Object03(c *gin.Context) {
	obj := &User{
		Name: "我是 User 对象返回",
		Age:  18,
	}
	c.JSON(http.StatusOK, obj)
}

func Jsonp(c *gin.Context) {
	obj := &User{
		Name: "我是 User 对象返回",
		Age:  18,
	}
	c.JSONP(http.StatusOK, obj)
}

func Xml(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{
		"name": "是我 xml",
	})
}

func Html(c *gin.Context) {
	// 需要配置 otherRouter.LoadHTMLGlob("./template/**")
	c.HTML(http.StatusOK, "index-01.html", gin.H{
		"title": "我是后台的数据",
	})
}

func GetParams(c *gin.Context) {
	pageNum := c.Query("pageNum")
	page := c.DefaultQuery("page", "1")

	c.JSON(http.StatusOK, gin.H{
		"page":    page,
		"pageNum": pageNum,
	})
}

func PostParams(c *gin.Context) {
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
}

func GetStruct(c *gin.Context) {
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
}
