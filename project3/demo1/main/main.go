package main

/**
字符串
  go中的字符串是由单个字节连接起来的   go语言的字符串的字节使用utf-8编码标识的unicode文本
*/
func main() {

	//1、 基本使用
	//var name string = "hello world"
	//fmt.Println(name)

	//2、注意事项
	//1、当字符串一旦赋值  将不能更改
	//var str string = "hello"
	//str[0] = 'e'    //编译不能通过

	//2、 字符串的两种表现形式
	//使用双引号 会识别转义字符
	//使用 反引号  ` 以字符串的原生形式进行输出，包括特殊字符

	//3、字符串的拼接   可以分行进行拼接，但是 + 号要在上一行  否则会编译不通过
	//var str1 string = "hello world "
	//str1 += "nihao"
	//fmt.Println(str1)

}
