package controller

import (
	"fmt"
	"github.com/axinger/ax-go-web/src/main/go/global"
	"github.com/axinger/ax-go-web/src/main/go/model/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BindController struct{}

func NewBindController() *BindController {
	return &BindController{}
}

// Build  请求参数绑定对象
func (i BindController) Build(r *global.Start) {

	// ShouldBindUri 参数绑定 uri方式 Login 需要配置 uri
	r.GET("/bind1/:user/:password", func(c *gin.Context) {
		var login dto.Login
		err := c.ShouldBindUri(&login)

		if err != nil {
			fmt.Printf("err = %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "参数错误" + err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "成功",
			"user":    login,
		})

	})

	// ShouldBind body方式 json
	r.POST("/bind2", func(c *gin.Context) {
		var login dto.Login
		err := c.ShouldBindJSON(&login)

		if err != nil {
			fmt.Printf("err = %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "参数错误" + err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "成功",
			"user":    login,
		})

	})

	// ShouldBindQuery,配置 form
	r.GET("/bind3", func(c *gin.Context) {
		var login dto.Login
		err := c.ShouldBindQuery(&login)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "参数错误" + err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "成功",
			"user":    login,
		})

	})
}
