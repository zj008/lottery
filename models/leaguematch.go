package models

import "github.com/astaxie/beego/orm"

type LeagueMatch struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type ExpertLeagueMatch struct {
	Id int `json:"id"`
	ExpertId int `json:"expert_id"`
	LeagueMatchId int `json:"league_match_id"`
}

func (l *LeagueMatch) TableName() string {
	return "leaguematch"
}

func (e *ExpertLeagueMatch) TableName() string {
	return "expert_leaguematch"
}

func init()  {
	orm.RegisterModel(new(LeagueMatch))
	orm.RegisterModel(new(ExpertLeagueMatch))
}

func GetExpertLeagueMatch(id int) (l []LeagueMatch, err error)  {
	o := orm.NewOrm()
	_, err = o.Raw("select * from leaguematch where id in" +
		" (select leaguematch_id from expert_leaguematch where expert_id = ?)", id).QueryRows(&l)
	return
}
