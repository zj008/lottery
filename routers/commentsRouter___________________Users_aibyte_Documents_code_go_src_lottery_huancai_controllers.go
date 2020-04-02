package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["lottery/huancai/controllers:ArticlesControllers"] = append(beego.GlobalControllerRouter["lottery/huancai/controllers:ArticlesControllers"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lottery/huancai/controllers:ArticlesControllers"] = append(beego.GlobalControllerRouter["lottery/huancai/controllers:ArticlesControllers"],
        beego.ControllerComments{
            Method: "GetDetail",
            Router: `/detail`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lottery/huancai/controllers:ExpertController"] = append(beego.GlobalControllerRouter["lottery/huancai/controllers:ExpertController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lottery/huancai/controllers:ExpertController"] = append(beego.GlobalControllerRouter["lottery/huancai/controllers:ExpertController"],
        beego.ControllerComments{
            Method: "GetHot",
            Router: `/hot`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lottery/huancai/controllers:MatchControllers"] = append(beego.GlobalControllerRouter["lottery/huancai/controllers:MatchControllers"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/foot`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lottery/huancai/controllers:NewsControllers"] = append(beego.GlobalControllerRouter["lottery/huancai/controllers:NewsControllers"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
