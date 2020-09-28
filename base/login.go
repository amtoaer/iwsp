package base

import (
	"iwsp/utils"

	"github.com/neucn/neugo"
)

// Login 登陆一网通/webVPN
func (s *Session) Login(username, password string, webVPN bool) {
	utils.Log("正在登陆中...")
	var platform neugo.Platform
	if webVPN {
		platform = neugo.WebVPN
		s.createURL = neugo.EncryptWebVPNUrl(s.createURL)
		s.infoURL = neugo.EncryptWebVPNUrl(s.infoURL)
		s.orderListURL = neugo.EncryptWebVPNUrl(s.orderListURL)
	} else {
		platform = neugo.CAS
	}
	client := neugo.NewSession()
	if err := neugo.Use(client).WithAuth(username, password).On(platform).Login(); err != nil {
		utils.Fatal(err)
	}
	s.client = client
	utils.Log("登陆成功")
}
