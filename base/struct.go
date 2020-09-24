package base

import "net/http"

// Session api/client/data的结构体
type Session struct {
	createURL string
	client    *http.Client
	data      interface{}
}

type fycc struct {
	ruleID     string
	bookDate   string
	periodName string
	bookCount  int
}
