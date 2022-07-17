package valid

import (
	"github.com/go-playground/validator/v10"
)

func StrMinLength(fl validator.FieldLevel) bool {
	// 先判断类型
	if str, ok := fl.Field().Interface().(string); ok {
		if len(str) < 6 {
			return false
		}
	}
	return true
}

func init() {

}
