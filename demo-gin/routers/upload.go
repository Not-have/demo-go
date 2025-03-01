package routers

import (
	"fmt"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

func Upload(r *gin.Engine) {
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	otherRouter := r.Group("/file", initMiddlewareOtherOne)

	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	otherRouter.POST("/upload", func(c *gin.Context) {
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
}
