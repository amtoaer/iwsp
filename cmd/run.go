package cmd

import (
	"flag"
	"fmt"
	"io/ioutil"
	"iwsp/base"
	"iwsp/utils"
	"os"
	"strings"
	"time"

	"github.com/mitchellh/go-homedir"
)

var (
	username  string
	password  string
	save      bool
	orderList bool
	webvpn    bool
	location  string
	duration  string
	debug     bool
)

func init() {
	flag.StringVar(&username, "u", "", "用户名")
	flag.StringVar(&password, "p", "", "密码")
	flag.BoolVar(&save, "s", false, "保存帐号密码到配置文件")
	flag.BoolVar(&orderList, "o", false, "输出历史预约信息")
	flag.BoolVar(&webvpn, "v", false, "使用webvpn")
	flag.StringVar(&location, "l", "", "预约地点")
	flag.StringVar(&duration, "t", "", "预约时段")
	flag.BoolVar(&utils.Debug, "d", false, "开启debug模式")
	flag.Usage = usage
}

// Run 用来承载主程序逻辑
func Run() {
	// 得到配置文件路径
	configPath, _ := homedir.Expand("~/.iwsp")
	// 未指定用户名或密码的情况下，尝试读取配置文件
	if len(username) == 0 || len(password) == 0 {
		content, err := ioutil.ReadFile(configPath)
		if err != nil {
			utils.Fatal("必须输入用户名和密码！\n可输入iwsp --help查看帮助信息。")
		}
		tmp := strings.Split(string(content), "\n")
		username = tmp[0]
		password = tmp[1]
		// 输入用户名密码的情况下，如果指定-s则覆盖写入帐号密码
	} else {
		if save {
			ioutil.WriteFile(configPath, []byte(username+"\n"+password), os.FileMode(0644))
		}
	}
	// 使用neugo进行登陆
	session := new(base.Session)
	session.Login(username, password, webvpn)
	// 如果指定-o则输出预约历史
	if orderList {
		session.GetOrderList()
		return
	}
	// 使用InitData函数得到需要post的数据以及获取信息的url
	session.InitData(location)
	// 设置需要post的数据（待修改）
	session.GetData().Set(13, time.Now().Format("2006-01-02"), duration, 1)
	// 发送预约请求
	session.Post()
}

func usage() {
	fmt.Println(`iwsp 东北大学场馆预约工具

	-u 一网通学号
	-p 一网通密码
	-s 保存学号密码到配置文件
	-v 使用webVPN，默认不使用
	-o 输出历史预约列表
	-l 预约地点，可选值fycc
	-t 预约时段，可选值
		07：00-10：00
		10：40-12：30
		12：30-14：00
		14：00-16：00
		16：00-18：00
		18：00-19：30
		19：30-21：00
	-d 启用debug模式
	-h 打印该帮助信息
	`)
}
