package main

import "fmt"

/**
  小数类型就是存放小数的
  小数类型分为单精度 float32 占用4个字节   float64 占用8个字节
  区别在表示的范围和精度不一样 浮点数容易丢失精度
*/
func main() {
	//容易丢失精度，应该选用高精度的float64  来保存更高精度
	var num1 float32 = -123.0000091
	var num2 float64 = -123.00000009001
	fmt.Println("num1:", num1, "num2:", num2)

	//go语言中 默认的浮点型是float64
	var num3 = 1.23
	fmt.Printf("num3的数据类型 %T", num3)
}
