package admin

import (
	"demo-gin-mvc/models"

	"github.com/gin-gonic/gin"
)

type UserControllers struct{}

func (con UserControllers) Index(c *gin.Context) {
	// 使用中间键的数据
	users := []models.User{} // 使用复数命名更清晰
	// users := models.User{} // 返回的是对象
	result := models.DB.Find(&users) // 连接查询数据

	// 传入指针，并处理错误
	// result := models.DB.Find(&users)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data": users,
	})
}
