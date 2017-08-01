/*
 提供的常用库，有一些常用的方法，方便使用
*/
package pack_fmt

import (
	"fmt"
)

var Pack1IntDoc int = 42
var PackFloatDoc = 3.14

//打印日志
func PrintLog_doc(str string) {
	fmt.Println("a.go is package mufunc.")
	fmt.Println(str)
}

//返回测试字符串
func ReturnStr_doc() string {
	return "hello main"
}
