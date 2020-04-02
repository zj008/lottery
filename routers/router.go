package routers

import (
	"lottery/huancai/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("*", &controllers.MainController{})
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/expert",
			beego.NSInclude(
				&controllers.ExpertController{},
			),
		),

		beego.NSNamespace("/articles",
			beego.NSInclude(
				&controllers.ArticlesControllers{},
			),
		),

		beego.NSNamespace("/match",
			beego.NSInclude(
				&controllers.MatchControllers{},
			),
		),

		beego.NSNamespace("/news",
			beego.NSInclude(
				&controllers.NewsControllers{},
			),
		),
	)

	beego.AddNamespace(ns)
}
