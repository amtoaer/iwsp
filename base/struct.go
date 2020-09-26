package base

import (
	"iwsp/utils"
	"net/http"
)

// Session api/client/data的结构体
type Session struct {
	createURL string
	infoURL   string
	client    *http.Client
	data      PostContent
	// 用于标记某个时段的人数
	countMap map[string]int
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
	Check(map[string]int)
}

// 风雨操场的set接口实现
func (f *fycc) Set(ruleID int, bookDate, periodName string, bookCount int) {
	f.RuleID = ruleID
	f.BookDate = bookDate
	f.PeriodName = periodName
	f.BookCount = bookCount
}

func (f *fycc) Check(m map[string]int) {
	value, ok := m[f.PeriodName]
	if !ok || value == 0 {
		utils.Fatal("时段不正确或预约人数已满")
	}
}
