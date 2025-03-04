package main

import (
	"demo-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/list", func(c *gin.Context) {
		users := []models.User{} // 使用复数命名更清晰
		// users := models.User{} // 使用复数命名更清晰

		// result := models.DB.Where("id>1 AND id<4").Find(&users) // 条件

		result := models.DB.Where("age LIKE?", "%22%").Find(&users) // 条件

		// 传入指针，并处理错误
		// result := models.DB.Find(&users)
		if result.Error != nil {
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}

		// 返回查询结果
		c.JSON(200, gin.H{
			"data": users,
		})
	})

	r.Run()
}
