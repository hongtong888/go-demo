package main

import "fmt"

//变量的介绍
func main() {

	//变量声明
	//1、指定变量类型，若不赋值，将输出默认值
	//var i int
	//i = 10
	//fmt.Println(i)

	//2、根据值 自行推导变量类型
	//var i = 10
	//fmt.Println(i)

	//3、 省略var关键字  使用 := 进行声明，但是左侧的变量不应该是已经声明过的变量 否则编译会报错
	i := 10
	fmt.Println(i)

}
