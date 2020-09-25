package base

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"iwsp/utils"
	"net/http"
)

// GetInfo 得到姓名与场次信息（待具体实现）
func (s *Session) GetInfo() (content string, err error) {
	var resp *http.Response
	if resp, err = s.client.Get(s.infoURL); err != nil {
		return
	}
	if content, err = readBody(resp); err != nil {
		return
	}
	return content, err
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
