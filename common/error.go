package common

import (
	"fmt"
	"os"
)

// 错误处理
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error： ", err.Error())
	}
}

// 日志处理
func Log(log string) {
	fmt.Println(log)
}
