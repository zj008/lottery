package controllers

import (
	"github.com/astaxie/beego"
	"lottery/huancai/lib"
	"lottery/huancai/models"
)

type ExpertController struct {
	beego.Controller
}

var host = beego.AppConfig.String("host")


// GetAll ...
// @Title Get All
// @Description get Channel
// @Param	type	query	string	false	"type of expert. Must be foot or basket"
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Channel
// @Failure 403
// @router / [get]
func (e *ExpertController) GetAll() {
	t := e.GetString("type")
	if t != "foot" && t != "basket" {
		lib.ResponseError(&e.Controller, lib.ErrorParam, nil)
	}
	var limit int64 = 100
	var offset int64 = 0
	if v, err := e.GetInt64("limit");err==nil{
		limit = v
	}
	if v, err := e.GetInt64("offset");err==nil{
		offset = v
	}
	n, experts, err := models.GetAllExpert(t, limit, offset)
	if err != nil{
		lib.ResponseError(&e.Controller, lib.ErrorGetAllExpert, err)
	}
	for i:=0;i<len(experts);i++{
		experts[i].Avatar = host + "/img/avatar/" + experts[i].Avatar
	}
	res := struct {
		Total int64 `json:"total"`
		Data interface{} `json:"data"`
	}{
		Total:n,
		Data:experts,
	}
	lib.ResponseSuccess(&e.Controller, res)
}

// GetHotExpert ...
// @Title Get Hot Exp
// @Description get Channel
// @Param	type	query	int 	false	"type of expert. 0:盈利榜， 1:人气榜"
// @Success 200 {object} models.Channel
// @Failure 403
// @router /hot [get]
func (e *ExpertController) GetHot() {
	t, err := e.GetInt64("type")
	if err != nil{
		lib.ResponseError(&e.Controller, lib.ErrorParam, err)
	}
	experts, err := models.GetHotExpert(t)
	if err != nil{
		lib.ResponseError(&e.Controller, lib.ErrorGetHotExpert, err)
	}
	for i:=0;i<len(experts);i++{
		experts[i].Avatar = host + "/img/avatar/" + experts[i].Avatar
	}
	lib.ResponseSuccess(&e.Controller, experts)
}

// GetHotExpert ...
// @Title Get Hot Exp
// @Description get Channel
// @Param	id	query	int 	false	"expert id"
// @Success 200 {object} models.Channel
// @Failure 403
// @router /detail [get]
func (e *ExpertController) GetDetail() {
	id, err := e.GetInt("id")
	if err != nil{
		lib.ResponseError(&e.Controller, lib.ErrorParam, err)
		return
	}
	expert, err := models.GetExpertDetail(id)
	if err != nil{
		lib.ResponseError(&e.Controller, lib.ErrorGetExpertDetail, err)
	}
	expert.Avatar = host + "/img/avatar/" + expert.Avatar
	lib.ResponseSuccess(&e.Controller, expert)
}

