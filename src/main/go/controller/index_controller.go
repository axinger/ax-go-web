package controller

import (
	"github.com/axinger/ax-go-web/src/main/go/global"
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

	// HTTP重定向
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com/")
	})

}

func (i IndexController) GetIndex() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"result": "index ok",
		})
	}
}
