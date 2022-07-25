package controller

import (
	"fmt"
	"github.com/axinger/ax-go-web/src/main/go/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FileController struct{}

func NewFileController() *FileController {
	return &FileController{}
}

// Build 文件控制器
func (i FileController) Build(r *global.Start) {

	// ShouldBindUri 参数绑定 uri方式 Login 需要配置 uri
	r.POST("/upload", func(c *gin.Context) {
		/// 多文件
		form, err2 := c.MultipartForm()
		if err2 != nil {
			return
		}
		_ = form.File["file"]

		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		name := c.PostForm("name")
		fmt.Print("name = ", name)
		c.SaveUploadedFile(file, "./"+file.Filename)
	})
}
