package base

import (
	"encoding/json"
	"io/ioutil"
	"iwsp/utils"
	"net/http"
	"strings"
)

// Get get某个地址并返回body和错误信息（仅用于测试）
func (s *Session) Get(url string) (content string, err error) {
	var resp *http.Response
	if resp, err = s.client.Get(url); err != nil {
		return
	}
	if content, err = readBody(resp); err != nil {
		return
	}
	return content, err
}

// Post post预约信息
func (s *Session) Post() {
	result, err := json.Marshal(s.data)
	utils.Log("序列化预约信息：", s.data)
	if err != nil {
		utils.Fatal("data序列化失败，请重试")
	}
	resp, err := s.client.Post(s.createURL, "application/json;charset=UTF-8", strings.NewReader(string(result)))
	if err != nil {
		utils.Fatal("post预约请求失败，请重试")
	}
	defer resp.Body.Close()
	utils.Log("预约返回结果：\n", resp.Body)
}

func readBody(resp *http.Response) (string, error) {
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	_ = resp.Body.Close()
	return string(content), nil
}
