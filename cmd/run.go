package cmd

import (
	"flag"
	"fmt"
	"os"
)

var (
	username string
	password string
	webvpn   bool
	location string
	time     string
)

func init() {
	flag.StringVar(&username, "u", "", "用户名")
	flag.StringVar(&password, "p", "", "密码")
	flag.BoolVar(&webvpn, "v", false, "是否使用webvpn")
	flag.StringVar(&location, "l", "", "预约地点")
	flag.StringVar(&time, "t", "", "预约时段")
	flag.Usage = usage
}

// Run is used to handle args
func Run() {
	if len(username) == 0 || len(password) == 0 {
		fatal("必须输入用户名和密码！\n可输入iwsp --help查看帮助信息。")
	}
}

func fatal(message string) {
	_, _ = fmt.Fprintln(os.Stderr, message)
	os.Exit(1)
}

func usage() {
	fmt.Println(`iwsp 东北大学场馆预约工具

	-u 一网通学号
	-p 一网通密码
	-v 使用webVPN，默认不使用
	-l 预约地点，可选值fycc
	-t 预约时段，可选值...`)
}
