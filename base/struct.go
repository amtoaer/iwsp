package base

import "net/http"

// Session api/client/data的结构体
type Session struct {
	createURL string
	client    *http.Client
	data      PostContent
}

type fycc struct {
	ruleID     string
	bookDate   string
	periodName string
	bookCount  int
}

// PostContent 所有post的结构体需要实现的接口
type PostContent interface {
	Set(string, string, string, int)
}

// 风雨操场的set接口实现
func (f *fycc) Set(ruleID, bookDate, periodName string, bookCount int) {
	f.ruleID = ruleID
	f.bookDate = bookDate
	f.periodName = periodName
	f.bookCount = bookCount
}
