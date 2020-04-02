package models

import "github.com/astaxie/beego/orm"

type Team struct {
	Id int `json:"id"`
	TeamType int `json:"team_type" description:"0:足球队, 1:篮球队"`
	FullName string `json:"full_name"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

func (t *Team) TableName() string {
	return "team"
}

func init() {
	orm.RegisterModel(new(Team))
}

func GetTeamById(id int) (t Team, err error) {
	o := orm.NewOrm()
	t = Team{Id:id}
	err = o.Read(&t)
	return
}