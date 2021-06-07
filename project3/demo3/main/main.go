package main

import (
	"fmt"
	"strconv"
)

/**
string类型转基本数据类型
*/
func main() {

	//1、使用fmt.Sprintf 方法进行转换
	//var num1 int = 90
	//str := fmt.Sprintf("%d", num1)
	//fmt.Printf("str type %T str %q",str,str)

	//2、 使用strconv包下
	//var num2 float64 = 123
	//str := strconv.FormatFloat(num2, 'f', 10, 64)
	//fmt.Printf("str type %T str %q",str,str)

	var num2 int = 123
	str := strconv.Itoa(num2)
	fmt.Printf("str type %T str %q", str, str)
}
