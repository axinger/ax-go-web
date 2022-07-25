package controller

import (
	"fmt"
	"github.com/axinger/ax-go-web/src/main/go/global"
	"github.com/axinger/ax-go-web/src/main/go/model/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (u UserController) Build(r *global.Start) {
	r.Handle("GET", "/userController", u.GetList())

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
}

func (u UserController) GetList() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"result": "index ok",
		})
	}
}
