package main

import (
	"fmt"
	"unsafe"
)

/**
go 中整数类型有int int8 int16 int32 int64  后面的数据代表位数
  int的位数根据当前的操作系统  32位和int32位等价

  unit 是无符号的
  rune 是有符号的  与int32类似 表示一个unicode
*/
func main() {
	//var num int64
	//num = 123
	//fmt.Printf("值：%v  类型 %T",num,num)

	var num = 123
	fmt.Printf("值： %v 。 类型 %T 。  占用的字节是 %d ", num, num, unsafe.Sizeof(num))
}
