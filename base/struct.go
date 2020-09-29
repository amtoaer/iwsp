package base

import (
	"iwsp/utils"
	"net/http"
)

// Session api/client/data的结构体
type Session struct {
	createURL    string
	cancelURL    string
	infoURL      string
	orderListURL string
	client       *http.Client
	data         PostContent
	// 用于标记某个时段的人数
	countMap map[string]int
}

// 取消预约需要的数据
type cancel struct {
	OrderID string `json:"orderId"`
}

// 风雨操场的post数据
type fycc struct {
	RuleID     int    `json:"ruleId"`
	BookDate   string `json:"bookDate"`
	PeriodName string `json:"periodName"`
	BookCount  int    `json:"bookCount"`
}

// 刘长春体育馆的post数据
type lcc struct {
	PeriodName string `json:"periodName"`
	BookCount  int    `json:"bookCount"`
	BookDate   string `json:"bookDate"`
	DeviceID   int    `json:"deviceId"`
	DeviceName string `json:"deviceName"`
	RuleID     int    `json:"ruleId"`
	SN         string `json:"sn"`
}

// 羽乒馆需要的“同行人”结构体
type user struct {
	Name string `json:"name"`
	No   string `json:"no"`
}

// 羽乒馆的post数据
type ypg struct {
	PeroidName string `json:"periodName"`
	BookCount  int    `json:"bookCount"`
	BookDate   string `json:"bookDate"`
	DeviceID   int    `json:"deviceId"`
	DeviceName string `json:"deviceName"`
	RuleID     int    `json:"ruleId"`
	UserList   []user `json:"userList"`
	SN         string `json:"sn"`
}

// post的返回结果
type returnData struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Data    string `json:"data"`
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
	utils.Log("通过时段/人数检查")
}
