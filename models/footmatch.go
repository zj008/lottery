package models

import "github.com/astaxie/beego/orm"


type FootMatchInfo struct {
	Id              int    `json:"id"`
	InfoId          int    `json:"info_id"`
	MatchStatus     int    `json:"match_status"`
	CardHighlight   int    `json:"card_highlight"`
	CornerKick      string `json:"corner_kick"`
	GuestRedCard    int    `json:"guest_red_card"`
	GuestYellowCard int    `json:"guest_yellow_card"`
	GuestScore      int    `json:"guest_score"`
	HalfScore       string `json:"half_score"`
	HomeRedCard     int    `json:"home_red_card"`
	HomeYellowCard  int    `json:"home_yellow_card"`
	HomeScore       int    `json:"home_score"`
	LeagueId        int    `json:"league_id"`
	LeagueName      string `json:"league_name"`
	MatchTime       string `orm:"column(match_time)" json:"match_time"`
	MatchId         int    `json:"match_id"`
	HomeName        string `json:"home_name"`
	HomeIcon        string `json:"home_icon"`
	GuestName       string `json:"guest_name"`
	GuestIcon       string `json:"guest_icon"`
}

func (f *FootMatchInfo) TableName() string {
	return "footmatch_info"
}

func init() {
	orm.RegisterModel(new(FootMatchInfo))
}

func GetFootballMatch(start string, end string) (n int64, d []FootMatchInfo, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("footmatch_info")
	n, err = qs.Filter("match_time__gte", start).Filter("match_time__lt", end).All(&d)
	if err != nil {
		return 0, d, err
	}
	return
}

