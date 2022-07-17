package controller

import (
	"fmt"
	"github.com/axinger/ax-go-web/src/main/go/global"
	"github.com/axinger/ax-go-web/src/main/go/model/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct{}

func NewIndexController() *IndexController {
	return &IndexController{}
}

func (i IndexController) Build(r *global.Start) {

	r.GET("/", func(c *gin.Context) {
		//c.JSON(http.StatusOK, gin.H{
		//	"message": "首页",
		//})

		h := gin.H{
			"msg":   "首页",
			"title": "我是go web",
			"list": gin.H{
				"name": "jim",
			},
		}

		c.HTML(http.StatusOK, "login.html", h)

	})

	// 参数绑定
	r.GET("/user", func(c *gin.Context) {

		var login dto.Login
		// get 用 ShouldBind
		// post 用 ShouldBindJSON
		// bind
		err := c.ShouldBind(&login)

		if err != nil {
			fmt.Printf("err = %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "参数错误",
			})
			return
		}

		fmt.Printf("login = %v\n", login)
		fmt.Println("Println login = ", login)
		c.JSON(http.StatusOK, gin.H{
			"message": "login",
			"user":    login,
		})

	})

	// HTTP重定向
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com/")
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
	{
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
}

func (i IndexController) GetIndex() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"result": "index ok",
		})
	}
}
