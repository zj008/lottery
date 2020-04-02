package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
)

type Expert struct {
	Id int	`orm:"column(id);auto;pk" json:"id"`
	Avatar string `orm:"column(avatar);size(80)" json:"avatar"`
	AvgOdds float64 `orm:"column(avg_odds)" json:"avg_odds"`
	BallRate string `orm:"column(ball_rate)" json:"ball_rate"`
	NickName string `orm:"column(nickname)" json:"nickname"`
	Slogan string `orm:"column(slogan)" json:"slogan"`
	HitRate float64 `orm:"column(hit_rate)" json:"hit_rate"`
	Foot int `orm:"column(foot)" json:"foot"`
	Basket int `orm:"column(basket)" json:"basket"`
}

type HotExpert struct {
	Id int `orm:"column(id)" json:"id"`
	Type int `orm:"column(type)" json:"type"`
	TopIndex int `orm:"column(top_index)" json:"top_index"`
	Popularity int `orm:"column(popularity)" json:"popularity"`
	EarningRate int `orm:"column(earning_rate)" json:"earning_rate"`
}

type HotExp struct {
	Id int	`orm:"column(id);auto;pk" json:"id"`
	Avatar string `orm:"column(avatar);size(80)" json:"avatar"`
	AvgOdds float64 `orm:"column(avg_odds)" json:"avg_odds"`
	BallRate string `orm:"column(ball_rate)" json:"ball_rate"`
	NickName string `orm:"column(nickname)" json:"nickname"`
	Slogan string `orm:"column(slogan)" json:"slogan"`
	HitRate float64 `orm:"column(hit_rate)" json:"hit_rate"`
	Foot int `orm:"column(foot)" json:"foot"`
	Basket int `orm:"column(basket)" json:"basket"`
	TopIndex int `orm:"column(top_index)" json:"top_index"`
	Popularity int `orm:"column(popularity)" json:"popularity"`
	EarningRate int `orm:"column(earning_rate)" json:"earning_rate"`
}

func (h *HotExpert) TableName() string {
	return "hot_expert"
}

func (e *Expert) TableName() string {
	return "expert"
}

func init() {
	orm.RegisterModel(new(Expert))
	orm.RegisterModel(new(HotExpert))
}

func GetAllExpert(t string, limit int64, offset int64) (e []Expert, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Expert))
	n, err := qs.Filter(t, 1).Limit(limit, offset).All(&e)
	if err != nil{
		return e, err
	}
	if n == 0{
		err = errors.New("no row found")
		return e, err
	}
	return e, nil
}

func GetHotExpert(t int64) (e []HotExp, err error) {
	o := orm.NewOrm()
	n, err := o.Raw("select a.id, a.top_index, a.popularity, a.earning_rate, b.foot, b.basket, b.avatar, b.avg_odds, b.ball_rate, b.nickname, b.slogan," +
		" b.hit_rate from hot_expert a left join expert b on a.id = b.id where a.type = ?", t).QueryRows(&e)
	if err != nil{
		return e, err
	}
	if n == 0 {
		err = errors.New("no row fond")
		return e, err
	}
	return e, nil
}