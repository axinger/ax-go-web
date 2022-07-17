package controller

import (
	"github.com/axinger/ax-go-web/src/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (u UserController) Build(g *global.Start) {
	g.Handle("GET", "/userController", u.GetList())
}

func (u UserController) GetList() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"result": "index ok",
		})
	}
}
