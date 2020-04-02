package models

import "github.com/astaxie/beego/orm"

type BasketballMatch struct {
	Id int `orm:"column(id)" json:"id"`
	GuestScore int `orm:"column(guest_score)" json:"guest_score"`
	HomeScore int `json:"home_score"`
	MatchId int `json:"match_id"`
	InfoId int `json:"info_id"`
	MatchStatus int `json:"match_status"`
	Status string `json:"status"`
	GuestId int `json:"guest_id"`
	HomeId int `json:"home_id"`
	MatchTime string `json:"match_time"`
	LeagueId int `json:"league_id"`
	LeagueName string `json:"league_name"`
	GuestName string `json:"guest_name"`
	GuestIcon string `json:"guest_icon"`
	HomeName string `json:"home_name"`
	HomeIcon string `json:"home_icon"`
}

func (b *BasketballMatch) TableName() string {
	return "basketmatch_info"
}

func init() {
	orm.RegisterModel(new(BasketballMatch))
}

func GetBasketMatch(s string, e string) (n int64, b []BasketballMatch, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(BasketballMatch))
	n, err = qs.Filter("match_time__lte", e).Filter("match_time__gte", s).Count()
	if err != nil{
		return 0, b, err
	}
	_, err = qs.Filter("match_time__lte", e).Filter("match_time__gte", s).All(&b)
	if err != nil{
		return 0, b, err
	}
	return
}


