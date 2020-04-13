package lib

import "fmt"

const (
	Success    = 0
	ErrorParam = 1
	ErrorJson  = 2

	ErrorGetNews    = 10
	ErrorUpdateNews = 11
	ErrorDeleteNews = 12

	ErrorGetExperts   = 20
	ErrorUpdateExpert = 21
	ErrorDeleteExpert = 22
	ErrorGetAllExpert = 23
	ErrorGetHotExpert = 24
	ErrorGetExpertDetail = 25



	ErrorGetArticles      = 30
	ErrorGetArticleDetail = 31
	ErrorUpdateArticle    = 32
	ErrorDeleteArticles   = 33

	ErrorTokenNull        = 40
	ErrorTokenExpired     = 41
	ErrorTokenNotValidYet = 42
	ErrorTokenMalformed   = 43
	ErrorTokenInvalid     = 44

	ErrorLogin = 50

	ErrorGetFootMatch = 60
	ErrorGetBasketMatch = 61

)

var m = map[int]string{
	Success:               "成功",
	ErrorParam:            "参数错误",
	ErrorJson:             "json数据错误",

	ErrorGetNews:          "获取资讯信息失败",
	ErrorUpdateNews:       "更新资讯信息失败",
	ErrorDeleteNews:       "删除资讯信息失败",

	ErrorGetExperts:       "获取专家信息谁败",
	ErrorUpdateExpert:     "更新专家信息失败",
	ErrorDeleteExpert:     "删除专家信息失败",
	ErrorGetAllExpert:     "获取专家列表失败",
	ErrorGetHotExpert:     "获取专家热榜失败",
	ErrorGetExpertDetail:  "获取专家详情失败",

	ErrorGetArticles:      "获取文章失败",
	ErrorGetArticleDetail: "获取文章详情失败",
	ErrorUpdateArticle:    "修改文章失败",
	ErrorDeleteArticles:   "删除文章失败",

	ErrorTokenNull:        "token不能为空",
	ErrorTokenExpired:     "Token is expired ",
	ErrorTokenNotValidYet: "Token not active yet ",
	ErrorTokenMalformed:   "That's not even a token ",
	ErrorTokenInvalid:     "Couldn't handle this token: ",

	ErrorLogin:            "登录失败",
}

func GetMsg(i int) string {
	return m[i]
}

type ErrorWithCode struct {
	Code int   `json:"code"`
	Err  error `json:"error"`
}

func GetError(code int) ErrorWithCode {
	return ErrorWithCode{code, fmt.Errorf(GetMsg(code))}
}

//优先使用
func GetErrorWithErr(code int, err error) ErrorWithCode {
	return ErrorWithCode{code, err}
}
