package routers

import (
	"lottery/huancai/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/*", &controllers.IndexController{})
	ns := beego.NewNamespace("/content",

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
