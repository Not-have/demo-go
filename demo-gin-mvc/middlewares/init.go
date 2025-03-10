package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func InitMiddleware(c *gin.Context) {
	//判断用户是否登录
	fmt.Println(time.Now())

	fmt.Println(c.Request.URL)

	c.Next() // 执行该中间件之后的逻辑
	//c.Abort() // 阻止执行该中间件之后的逻辑

	// 设置数据
	c.Set("request", "中间件设置的数据")

	//定义一个goroutine统计日志  当在中间件或 handler 中启动新的 goroutine 时，不能使用原始的上下文（c *gin.Context）， 必须使用其只读副本（c.Copy()）
	cCp := c.Copy()

	// 使用 goroutine 处理异步请求，也就是携程
	// 使用了 goroutine 后，不会阻塞请求，不能使用上下文的 c *gin.Context，只能使用 cCp := c.Copy() 的只读副本
	// 为了防止 goroutine 泄漏，确保在请求处理程序返回时或者使用 defer 语句
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Done! in path " + cCp.Request.URL.Path)
	}()
}
