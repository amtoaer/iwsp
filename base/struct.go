package base

import "net/http"

// Session api/client/data的结构体
type Session struct {
	createURL string
	infoURL   string
	client    *http.Client
	data      PostContent
}

type fycc struct {
	RuleID     int    `json:"ruleId"`
	BookDate   string `json:"bookDate"`
	PeriodName string `json:"periodName"`
	BookCount  int    `json:"bookCount"`
}

// PostContent 所有post的结构体需要实现的接口
type PostContent interface {
	Set(ruleID int, bookDate, periodName string, bookCount int)
}

// 风雨操场的set接口实现
func (f *fycc) Set(ruleID int, bookDate, periodName string, bookCount int) {
	f.RuleID = ruleID
	f.BookDate = bookDate
	f.PeriodName = periodName
	f.BookCount = bookCount
}
