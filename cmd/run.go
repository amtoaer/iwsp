package cmd

import (
	"flag"
	"fmt"
	"iwsp/base"
	"iwsp/utils"
	"time"
)

var (
	username  string
	password  string
	orderList bool
	webvpn    bool
	location  string
	duration  string
	debug     bool
)

func init() {
	flag.StringVar(&username, "u", "", "用户名")
	flag.StringVar(&password, "p", "", "密码")
	flag.BoolVar(&orderList, "o", false, "输出历史预约信息")
	flag.BoolVar(&webvpn, "v", false, "使用webvpn")
	flag.StringVar(&location, "l", "", "预约地点")
	flag.StringVar(&duration, "t", "", "预约时段")
	flag.BoolVar(&utils.Debug, "d", false, "开启debug模式")
	flag.Usage = usage
}

// Run 用来承载主程序逻辑
func Run() {
	if len(username) == 0 || len(password) == 0 {
		utils.Fatal("必须输入用户名和密码！\n可输入iwsp --help查看帮助信息。")
	}
	session := new(base.Session)
	session.Login(username, password, webvpn)
	if orderList {
		session.GetOrderList()
		return
	}
	session.InitData(location)
	session.GetData().Set(13, time.Now().Format("2006-01-02"), "16:00-18:00", 1)
	session.Post()
}

func usage() {
	fmt.Println(`iwsp 东北大学场馆预约工具

	-u 一网通学号
	-p 一网通密码
	-v 使用webVPN，默认不使用
	-o 输出历史预约列表
	-l 预约地点，可选值fycc
	-t 预约时段，可选值...
	-d 启用debug模式
	-h 打印该帮助信息
	`)
}
