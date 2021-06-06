package main

import "fmt"

/**
	goland中没有专门的字符类型，如果要存储单个字符 一般使用byte来保存
   字符串就是一串固定长度的字符连续连接起来的字符序列  	go的字符串不同是由字节组成
*/
func main() {
	//var c1 byte = 'a'
	//var c2 byte = 'b'
	//
	//fmt.Println("c1=",c1,"c2=",c2)  //输出的是字符的码值
	//fmt.Printf("c1=%c,c2=%c \n",c1,c2)  //格式化输出对应的字符
	//fmt.Printf("c1=%d,c2=%d",c1,c2)

	var c3 int = '北'
	fmt.Printf("c3=%c \n", c3) //格式化输出对应的字符
	fmt.Printf("c3=%d", c3)
}
