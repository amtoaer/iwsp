package base

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"iwsp/utils"
	"net/http"
	"regexp"
	"strings"
)

func (s *Session) get(url string) (content string, err error) {
	var resp *http.Response
	if resp, err = s.client.Get(url); err != nil {
		return
	}
	if content, err = readBody(resp); err != nil {
		return
	}
	return content, err
}

// GetOrderList 得到历史预约列表
func (s *Session) GetOrderList() {
	content, err := s.get("http://book.neu.edu.cn/booking/page/orderList")
	if err != nil {
		utils.Fatal(err)
	}
	r := regexp.MustCompile("var orderList = (.+);")
	matchResult := r.FindStringSubmatch(content)
	if len(matchResult) < 2 {
		utils.Fatal("正则表达式匹配失败！")
	}
	orderList := strings.ReplaceAll(matchResult[1], "'", "\"")
	var container []map[string]interface{}
	err = json.Unmarshal([]byte(orderList), &container)
	if err != nil {
		utils.Log(err)
		utils.Fatal("预约列表解析失败！")
	}
	utils.Log("预约列表解析成功！")
	output := func(m map[string]interface{}) {
		getStatus := func(status float64) string {
			switch status {
			case 0:
				return "已预约"
			case 1:
				return "已入场"
			case 2:
				return "已出场"
			case 3:
				return "已完成"
			case 4:
				return "已取消"
			case 5:
				return "已关闭"
			case 6:
				return "已过期"
			default:
				return "未知状态"
			}
		}
		var (
			position = m["ruleName"].(string)
			date     = m["bookPeriodStartTime"].(string)[0:10]
			duration = m["bookPeriodName"].(string)
			status   = getStatus(m["status"].(float64))
		)
		fmt.Printf("%14s%23s%23s%10s\n", position, date, duration, status)
	}
	fmt.Printf("%14s%20s%20s%14s\n", "地点", "日期", "时段", "状态")
	for _, value := range container {
		output(value)
	}
}

// Post post预约信息
func (s *Session) Post() {
	data, err := json.Marshal(s.data)
	utils.Log("序列化预约信息：", s.data)
	if err != nil {
		utils.Fatal("data序列化失败，请重试")
	}
	utils.Log("序列化结果:", string(data))
	utils.Log("开始进行post请求...")
	resp, err := s.client.Post(s.createURL, "application/json;charset=UTF-8", bytes.NewBuffer(data))
	if err != nil {
		utils.Fatal("post预约请求失败，请重试")
	}
	fmt.Println(resp.Status)
	result, err := readBody(resp)
	utils.Log("预约返回结果：\n", result)

}

func readBody(resp *http.Response) (string, error) {
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	_ = resp.Body.Close()
	return string(content), nil
}
