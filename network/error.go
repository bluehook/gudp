package network

import (
	"fmt"
	"os"
)

// 错误处理
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error： ", err.Error())
	}
}

// 日志处理
func log(log string) {
	fmt.Println(log)
}
