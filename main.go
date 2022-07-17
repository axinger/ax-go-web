package main

import (
	"fmt"
	"github.com/axinger/ax-go-web/src/main/go/config"
	"github.com/axinger/ax-go-web/src/main/go/controller"
	"github.com/axinger/ax-go-web/src/main/go/global"
	"github.com/axinger/ax-go-web/src/main/go/valid"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

func main() {

	// 参数验证规则, 放vaild init 有的顺序问题
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("strMinLength", valid.StrMinLength)
	}
	global.Ignite().
		Mount(
			controller.NewIndexController(),
			controller.NewUserController(),
			controller.NewTestController(),
			controller.NewBindController(),
		).
		Run(":8080")
}

var appConfig config.AppConfig

var mysqlConfig config.MysqlConfig

// 初始化配置文件
func init() {
	//viper.SetConfigName("resources/config")
	//viper.SetConfigType("yaml")
	////viper.AddConfigPath("/etc/appname/") // 查找配置文件所在路径
	////viper.AddConfigPath("$HOME/.appname") // 多次调用AddConfigPath，可以添加多个搜索路径
	//viper.AddConfigPath(".") // 还可以在工作目录中搜索配置文件
	//viper.AddConfigPath("./conf") // 还可以在工作目录中搜索配置文件

	//viper.SetConfigFile("resources/config.yml")

	viper.AddConfigPath("src/resources")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	//viper.AddConfigPath("resources/config.yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("读取config文件失败: = ", err)
	} else {
		if err := viper.Unmarshal(&appConfig); err != nil {
			fmt.Println("解析config文件失败: = ", err)
		} else {
			fmt.Println("解析config文件成功:: = ", appConfig)
		}
	}

	viper.AddConfigPath("src/resources")
	viper.SetConfigName("mysql")
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("读取mysql文件失败:: %v \r", err)
	} else {
		if err := viper.Unmarshal(&mysqlConfig); err != nil {
			fmt.Println("mysqlConfig =================== ", mysqlConfig)
		} else {
			fmt.Println("解析mysql文件成功:: = ", mysqlConfig)
		}

	}

}
