package base

// InitData 通过预约地点得到用于提交的PostContent
func (s *Session) InitData(location string) {
	s.orderListURL = "http://book.neu.edu.cn/booking/page/orderList"
	s.cancelURL = "http://book.neu.edu.cn/booking/order/cancel"
	switch location {
	case "fycc":
		s.data = &fycc{}
		s.infoURL = "http://book.neu.edu.cn/booking/page/rule/13"
		s.createURL = "http://book.neu.edu.cn/booking/order/create"
	}
}

// GetData 返回struct内的PostContent
func (s *Session) GetData() PostContent {
	return s.data
}

// IsDataEmpty 判断data是否为空（即判断地点参数是否正确设置）
func (s *Session) IsDataEmpty() bool {
	if s.data == nil {
		return true
	}
	return false
}
