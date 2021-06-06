package main

import "fmt"

/**
  指针：
    基础数据，变量存的是值，也称之为值类型   获取变量的地址   &
    指针类型  指针变量存的是地址，这个地址指向的空间才是存的值
*/
func main() {
	//1、获取基本数据类型变量地址
	var num = 123
	//fmt.Printf("num的地址是 %v",&num)

	//2、声明指针变量
	var ptr = &num
	fmt.Println("指针变量的值是ptr:", ptr)
	fmt.Printf("ptr %T \n", ptr)
	fmt.Println("指针变量ptr,指向的值是:", *ptr)

	//修改指针变量指向的值，原数值也会相应的进行修改
	*ptr = 100
	fmt.Println("num 的值是:", num)

}
