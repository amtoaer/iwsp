package utils

import (
	"fmt"
	"os"
)

var (
	// Debug 标记是否开启debug模式
	Debug = false
)

// Fatal 打印标准错误并退出程序
func Fatal(message interface{}) {
	_, _ = fmt.Fprintln(os.Stderr, message)
	os.Exit(1)
}

// Log 在debug模式下输出调试信息
func Log(message ...interface{}) {
	if Debug {
		fmt.Println(message...)
	}
}
