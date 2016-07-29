package home

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/open-falcon/fe/g"
)

func ConfigRoutes() {
	beego.Router("/", &HomeController{})

	beego.Get("/health", func(ctx *context.Context) {
		ctx.Output.Body([]byte("ok"))
	})

	// HEAD method may be used for health check such as aliyun SLB
	beego.Head("/health", func(ctx *context.Context) {
		ctx.Output.Body([]byte(""))
	})

	beego.Get("/version", func(ctx *context.Context) {
		ctx.Output.Body([]byte(g.VERSION))
	})
}
