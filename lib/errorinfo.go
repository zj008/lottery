package lib

const (
	Success    = 0
	ErrorParam = 1

	ErrorGetAllExpert = 10
	ErrorGetHotExpert = 11
	ErrorGetExpertDetail = 12

	ErrorGetArticles      = 20
	ErrorGetArticleDetail = 21

	ErrorGetFootMatch   = 30
	ErrorGetBasketMatch = 31

	ErrorGetNews = 40
)

var m = map[int]string{
	Success:               "ok",
	ErrorParam:            "参数错误",
	ErrorGetAllExpert:     "获取专家列表失败",
	ErrorGetHotExpert:     "获取排行榜专家失败",
	ErrorGetExpertDetail:  "获取专家详情失败",
	ErrorGetArticles:      "获取专家文章失败",
	ErrorGetArticleDetail: "获取文章内容失败",
	ErrorGetFootMatch:     "获取足球比赛信息失败",
	ErrorGetBasketMatch:   "获取篮球比赛信息失败",
	ErrorGetNews:          "获取资讯信息失败",
}

func GetMsg(i int) string {
	return m[i]
}
