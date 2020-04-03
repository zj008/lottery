package models

import "github.com/astaxie/beego/orm"

type News struct {
	Id         int    `json:"-"`
	NewsType   int    `json:"-"`
	CreateTime string `json:"create_time"`
	Source     string `json:"source"`
	Title      string `json:"title"`
	Url string 	`json:"-"`
	Docid string `json:"-"`
	Content string `json:"content"`
	Img string `json:"img"`
}

func (n *News) TableName() string {
	return "news"
}

func init() {
	orm.RegisterModel(new(News))
}

func GetNews(t int, limit int64, offset int64) (c int64, n []News, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(News))
	c, err = qs.Filter("news_type", t).Count()
	if err != nil{
		return c, n, err
	}
	_, err = qs.Filter("news_type", t).OrderBy("-create_time").Limit(limit, offset).All(&n)
	return
}