package models

import "github.com/astaxie/beego/orm"

type Matches struct {
	Id            int    `json:"id"`
	InfoId        int    `json:"info_id"`
	MatchStatus   int    `json:"match_status"`
	CardHighlight int    `json:"card_highlight"`
	GuestScore    int    `json:"guest_score"`
	CategoryId    int    `json:"category_id"`
	HomeScore     int    `json:"home_score"`
	LeagueId      int    `json:"league_id"`
	LeagueName    string `json:"league_name"`
	MatchTime     string `orm:"column(match_time)" json:"match_time"`
	HomeName      string `json:"home_name"`
	GuestName     string `json:"guest_name"`
}

type ArticleMatch struct {
	Id        int `json:"id"`
	ArticleId int `json:"article_id"`
	InfoId    int `json:"info_id"`
}

func (m *Matches) TableName() string {
	return "matches"
}

func (a *ArticleMatch) TableName() string {
	return "article_match"
}

func init() {
	orm.RegisterModel(new(Matches))
	orm.RegisterModel(new(ArticleMatch))
}

func GetMatchesByArticleId(id int) (m []Matches, err error) {
	o := orm.NewOrm()
	_, err = o.Raw("select * from matches where info_id in (select info_id from article_match where article_id = ?)", id).QueryRows(&m)
	return
}
