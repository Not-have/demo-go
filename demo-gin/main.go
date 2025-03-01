package main

import (
	"demo-gin/models"
	routers "demo-gin/routers"
	"fmt"
	"net/http"
	"path"
	"text/template"

	"github.com/gin-gonic/gin"
)

func initMiddlewareOne(c *gin.Context) {
	fmt.Println("我是全局中间件")

	// 终止请求
	// c.Abort()

	c.Next()

}

func main() {
	r := gin.Default()

	r.SetFuncMap(template.FuncMap{
		"UniToTime": models.UnixToTime,
	})

	r.LoadHTMLGlob("./template/**")

	r.Use(initMiddlewareOne)

	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.POST("/upload", func(c *gin.Context) {
		/*
			// // 单文件
			file, _ := c.FormFile("file")

			dst := path.Join("./upload", file.Filename)

			// 上传文件至指定的完整文件路径
			c.SaveUploadedFile(file, dst)

			c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
		*/
		// 多文件
		form, _ := c.MultipartForm()
		files := form.File["file[]"]

		for _, file := range files {

			// 上传文件至指定目录
			dst := path.Join("./upload", file.Filename)
			c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})

	// 路由分组
	routers.Base(r)
	routers.Other(r)

	r.Run(":8010") // listen and serve on 0.0.0.0:8080
}
