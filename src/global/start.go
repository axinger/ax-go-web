package global

import (
	"github.com/axinger/ax-go-web/src/filter"
	"github.com/gin-gonic/gin"
)

type Start struct {
	*gin.Engine
}

func Ignite() *Start {

	r := gin.Default()
	start := Start{r}

	// 注册一个全局中间件,拦截器
	r.Use(filter.StatCost(), filter.TokenHandler())
	r.LoadHTMLGlob("src/resources/templates/*")

	//router.Static第一个参数为路由匹配地址，第二个参数为静态资源相对路径
	r.Static("../static", "src/resources/static")

	return &start
}

// Mount 挂载需要执行的 gin.Handle
func (g *Start) Mount(controllers ...IController) *Start {
	for _, controller := range controllers {
		controller.Build(g)
	}
	return g
}

func (g *Start) Run(addr string) {
	if addr == "" {
		addr = ":8080"
	}
	g.Engine.Run(addr)
}
