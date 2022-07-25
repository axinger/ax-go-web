package filter

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// StatCost 是一个统计耗时请求耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {

		params := c.Params
		log.Println("StatCost params", params)
		log.Println("StatCost params-action", c.Param("action"))

		start := time.Now()
		c.Set("name", "小王子") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序
		c.Next()
		// 不调用该请求的剩余处理程序
		// c.Abort()
		// 计算耗时
		cost := time.Since(start)
		log.Println("cost:", cost)
	}
}

func TokenHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("token")
		log.Println("token:", token)
		log.Println("FullPath:", c.FullPath())
		// FullPath: /shop/cart
		log.Println("FullPath:", c.FullPath())

		log.Println("FullPath:", (c.FullPath() != "/" || c.FullPath() != "/index") && token == "")

		//if !(c.FullPath() == "/" || c.FullPath() == "/index") && token == "" {
		//	c.Abort()
		//	c.JSON(http.StatusOK, gin.H{
		//		"message": "未授权",
		//	})
		//	return
		//}

		params := c.Params
		log.Println("StatCost params", params)
		log.Println("StatCost params-action", c.Param("action"))

		start := time.Now()
		c.Set("name", "小王子") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序

		c.Next()
		// 不调用该请求的剩余处理程序

		// 计算耗时
		cost := time.Since(start)
		log.Println("cost:", cost)
	}
}
