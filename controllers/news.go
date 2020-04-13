package controllers

import (
	"github.com/astaxie/beego"
	"lottery/huancai/lib"
	"lottery/huancai/models"
	"strings"
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
	host := beego.AppConfig.String("host")
	for i:=0;i<len(news);i++{
		news[i].Content = strings.ReplaceAll(news[i].Content, "img/news", host+"/img/news")
		news[i].Img = host + "/img/news/" + news[i].Img
	}
	res := struct {
		Total int64 `json:"total"`
		Data interface{} `json:"data"`
	}{
		Total:c,
		Data:news,
	}
	lib.ResponseSuccess(&n.Controller, res)
}

