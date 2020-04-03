package models

import (
	"github.com/astaxie/beego/orm"
)

type Articles struct {
	Id int `orm:"column(id)" json:"id"`
	ExpertId int `orm:"column(expert_id)" json:"expert_id"`
	IsWin int `orm:"column(is_win)" json:"is_win"`
	LotteryCategoryId int `orm:"column(lottery_category_id)" json:"lottery_category_id"`
	LotteryCategoryName string `orm:"column(lottery_category_name)" json:"lottery_category_name"`
	Price float64 `orm:"column(price)" json:"price"`
	PublishTime string `orm:"column(publish_time)" json:"publish_time"`
	Title string `orm:"column(title)" json:"title"`
	MatchList []Matches `orm:"-" json:"match_list"`
}

type ArticleDetail struct {
	Id int `orm:"column(id)" json:"id"`
	Content string `orm:"column(content)" json:"content"`
	MatchList []Matches `orm:"-" json:"-"`
}

func (a *ArticleDetail) TableName() string {
	return "article_details"
}

func (a *Articles) TableName() string {
	return "articles"
}

func init()  {
	orm.RegisterModel(new(Articles))
	orm.RegisterModel(new(ArticleDetail))
}

func GetArticlesByUser(u int, limit int64, offset int64) (n int64, a []Articles, e error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Articles))
	n, err := qs.Filter("expert_id", u).Count()
	if err != nil{
		return 0, a, err
	}
	_, err = qs.Filter("expert_id", u).OrderBy("-publish_time").Limit(limit, offset).All(&a)
	if err != nil{
		return 0, a, err
	}

	for i:=0;i<len(a);i++{
		m, err := GetMatchesByArticleId(a[i].Id)
		if err != nil{
			continue
		}
		a[i].MatchList = m
	}
	return n, a, nil
}

func GetArticleDetail(id int) (d ArticleDetail, err error) {
	o := orm.NewOrm()
	d = ArticleDetail{Id:id}
	err = o.Read(&d)
	return
}
