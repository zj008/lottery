package controllers

import (
	"github.com/astaxie/beego"
	"lottery/huancai/lib"
	"lottery/huancai/models"
)

type ExpertController struct {
	beego.Controller
}

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
	experts, err := models.GetAllExpert(t, limit, offset)
	if err != nil{
		lib.ResponseError(&e.Controller, lib.ErrorGetAllExpert, err)
	}
	lib.ResponseSuccess(&e.Controller, experts)
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
	lib.ResponseSuccess(&e.Controller, experts)
}