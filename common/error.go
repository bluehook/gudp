package common

import (
	"fmt"
	"os"
)

//#网络错误
type NetError struct {
	os string
}

func (self *NetError) Error() string {
	return self.os
}

func NewNetError(os string) error {
	return &NetError{os}
}

//#错误处理
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error： ", err.Error())
	}
}

//#日志处理
func Log(log string) {
	fmt.Println(log)
}
