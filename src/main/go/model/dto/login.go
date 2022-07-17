package dto

type Login struct {
	User string `form:"user" json:"user" uri:"user" binding:"required"`
	//Password string `form:"password" json:"password" uri:"password" binding:"required"`

	// 自定义验证
	Password string `form:"password" json:"password" uri:"password" binding:"required,strMinLength"`
}
