package controller

import (
	"fmt"
	"github.com/axinger/ax-go-web/src/main/go/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TestController struct{}

func NewTestController() *TestController {
	return &TestController{}
}

func (t TestController) Build(r *global.Start) {

	r.GET("/test1", func(c *gin.Context) {

		params := c.Params
		fmt.Println("params", params)
		//c.JSON(http.StatusOK, gin.H{
		//	"message": "首页",
		//})

		c.String(http.StatusOK, "%v", "bb")
	})

	// restful 风格
	// *other 追加路径
	r.GET("/test2/:name/*other", func(c *gin.Context) {
		params := c.Params
		name, _ := params.Get("name")
		other, _ := params.Get("other")
		other1, _ := params.Get("other1")

		c.JSON(http.StatusOK, gin.H{
			"path":   "test2",
			"msg":    "restful 风格",
			"name":   name,
			"other":  other,
			"other1": other1,
		})
	})

	// url 风格
	r.GET("/test3", func(c *gin.Context) {
		// DefaultQuery 参数不存在,返回默认值
		// Query 返回空字串
		name := c.DefaultQuery("name", "")
		age := c.Query("age")

		fmt.Println(strconv.Atoi(age))

		fmt.Printf("name:%v,age:%v", name, age)

		c.JSON(http.StatusOK, gin.H{
			"message": "test3",
		})
	})

	// 路由重定向
	r.GET("/redirect1", func(c *gin.Context) {
		// 指定重定向的URL
		c.Request.URL.Path = "/redirect2"
		r.HandleContext(c)
	})
	r.GET("/redirect2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world"})
	})

	// 路由组
	shopGroup := r.Group("/shop")

	// http://localhost:8080/shop/index  要带前部分
	shopGroup.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "shop_index",
		})

	})
	shopGroup.GET("/cart", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "cart",
		})

	})
	shopGroup.POST("/checkout", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "checkout",
		})
	})

}

func (t TestController) GetList() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"result": "index ok",
		})
	}
}
