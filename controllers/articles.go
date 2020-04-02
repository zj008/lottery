package controllers

import (
	"github.com/astaxie/beego"
	"lottery/huancai/lib"
	"lottery/huancai/models"
)

type ArticlesControllers struct {
	beego.Controller
}

// GetAll ...
// @Title Get All
// @Description get Channel
// @Param	uid 	query	int 	false	"expert id. Must be an int"
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Channel
// @Failure 403
// @router / [get]
func (a *ArticlesControllers) Get() {
	u, err := a.GetInt("uid")
	if err != nil{
		lib.ResponseError(&a.Controller, lib.ErrorParam, err)
		return
	}
	var limit int64 = 10
	var offset int64 = 0
	if v, err := a.GetInt64("limit"); err == nil{
		limit = v
	}
	if v, err := a.GetInt64("offset"); err == nil{
		offset = v
	}
	n, articles, err := models.GetArticlesByUser(u, limit, offset)
	if err != nil{
		lib.ResponseError(&a.Controller, lib.ErrorGetArticles, err)
		return
	}
	data := struct {
		Total int64
		Data interface{}
	}{
		Total:n,
		Data:articles,
	}
	lib.ResponseSuccess(&a.Controller, data)
}

// GetAll ...
// @Title Get All
// @Description get article detail
// @Param   id 	query	int 	false	"article id. Must be an int"
// @Success 200 {object} models.Channel
// @Failure 403
// @router /detail [get]
func (a *ArticlesControllers) GetDetail() {
	id, err := a.GetInt("id")
	if err != nil{
		lib.ResponseError(&a.Controller, lib.ErrorParam, err)
		return
	}
	detail, err := models.GetArticleDetail(id)
	if err != nil{
		lib.ResponseError(&a.Controller, lib.ErrorGetArticleDetail, err)
		return
	}
	lib.ResponseSuccess(&a.Controller, detail)
}