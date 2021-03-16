package main

import "fmt"

func main() {
	//变量声明与定义
	//int 类型默认值为 0
	//变量定义格式01:   var  变量名 类型  声明
	var num int
	fmt.Println(num)
	num = 10
	fmt.Println(num)
	num += 10
	fmt.Println(num)

	//浮点型使用float  float32代表内存中32个字节存储    float64代表内存中64个字节存储
	//变量定义格式02 : var  变量名 类型 = 值 定义
	var long float64 = 2.2
	var width float64 = 2.2

	fmt.Println(long * width)

	//变量定义格式03 :  变量名 := 值  自动推导类型
	//自动推导类型 使用 变量:= 值 GO会根据所赋的值自动推导出变量的类型  开发中常用的形式
	pi := 3.1415926
	r := 2.5
	s := pi * r * r
	fmt.Println(s)

	price := 2.5
	num = 2
	//GO 语言中不同数据类型不能进行计算操作  需要进行类型转换
	fmt.Println(price * float64(num))

	a := false  //bool 布尔类型
	b := 10     //int 整型
	c := 3.14   //float64 浮点型
	d := 'a'    //byte 字节类型
	e := "瞅你咋地" //string 字符串类型

	fmt.Println(a, b, c, d, e)
}
