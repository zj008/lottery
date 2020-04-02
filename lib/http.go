package lib

import "github.com/astaxie/beego"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseSuccess(c *beego.Controller, data interface{}) {
	r := Response{
		Code: 0,
		Msg:  "success",
		Data: data,
	}
	c.Data["json"] = r
	c.ServeJSON()
}

func ResponseError(c *beego.Controller, code int, err error)  {
	msg := GetMsg(code)
	if err != nil{
		msg += err.Error()
	}
	r := Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
	c.Data["JSON"] = r
	c.ServeJSON()
}