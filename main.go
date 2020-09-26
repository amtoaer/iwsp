package main

import (
	"flag"
	"fmt"
	"iwsp/cmd"
)

var (
	v       bool
	version string
)

func init() {
	flag.BoolVar(&v, "version", false, "查看版本")
}

func main() {
	flag.Parse()
	if v {
		fmt.Println("当前版本：" + version)
		return
	}
	cmd.Run()
}
