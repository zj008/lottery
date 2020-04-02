package controllers

import (
	"github.com/astaxie/beego"
	"lottery/huancai/lib"
	"lottery/huancai/models"
)

type MatchControllers struct {
	beego.Controller
}

// GetAll ...
// @Title Get All
// @Description get article detail
// @Param   date 	query	string 	false	"match data. like 2020-03-29"
// @Success 200 {object} models.Channel
// @Failure 403
// @router /foot [get]
func (m *MatchControllers) Get() {
	date := m.GetString("date")
	start := date + " 00:00:00"
	end := date + " 23:59:59"
	n, f, err := models.GetFootballMatch(start, end)
	if err != nil{
		lib.ResponseError(&m.Controller, lib.ErrorGetFootMatch, err)
	}
	res := struct {
		Total int64
		Data interface{}
	}{
		Total:n,
		Data:f,
	}
	lib.ResponseSuccess(&m.Controller, res)
}


// GetAll ...
// @Title Get All
// @Description get article detail
// @Param   date 	query	string 	false	"match data. like 2020-03-29"
// @Success 200 {object} models.Channel
// @Failure 403
// @router /basket [get]
func (m *MatchControllers) GetBasket() {
	d := m.GetString("date")
	start := d + " 00:00:00"
	end := d + " 23:59:59"
	n, b, err := models.GetBasketMatch(start, end)
	if  err != nil{
		lib.ResponseError(&m.Controller, lib.ErrorGetBasketMatch, err)
		return
	}
	res := struct {
		Total int64
		Data interface{}
	}{
		Total:n,
		Data:b,
	}
	lib.ResponseSuccess(&m.Controller, res)
}

