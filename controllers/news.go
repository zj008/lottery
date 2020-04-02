package controllers

import (
	"github.com/astaxie/beego"
	"lottery/huancai/lib"
	"lottery/huancai/models"
)

type NewsControllers struct {
	beego.Controller
}

// GetAll ...
// @Title Get All
// @Description get article detail
// @Param   type 	query	string 	false	"0:足球, 1:篮球"
// @Param   limit 	query	string 	false	"limit"
// @Param   offset 	query	string 	false	"offset"
// @Success 200 {object} models.Channel
// @Failure 403
// @router / [get]
func (n *NewsControllers) Get() {
	t, err := n.GetInt("type")
	if err != nil{
		lib.ResponseError(&n.Controller, lib.ErrorParam, err)
	}
	var limit int64 = 10
	var offset int64 = 0
	if v, err := n.GetInt64("limit"); err == nil{
		limit = v
	}
	if v, err := n.GetInt64("offset"); err == nil{
		offset = v
	}
	c, news, err := models.GetNews(t, limit, offset)
	if err != nil{
		lib.ResponseError(&n.Controller, lib.ErrorGetNews, err)
	}
	res := struct {
		Total int64
		Data interface{}
	}{
		Total:c,
		Data:news,
	}
	lib.ResponseSuccess(&n.Controller, res)
}