package main

import "fmt"

func main() {
	a := 10
	b := 3.2
	//fmt.Println(a ..interface{})  可以输出任何格式 带换行
	fmt.Println(a)
	fmt.Println(b)
	//fmt.Println(a ..interface{})  可以输出任何格式 不带换行
	fmt.Print(a)
	fmt.Print(b)
	fmt.Print("\n")

	//%d  占位符，表示输出一个整形数据
	//%f  占位符，表示输出一个浮点型数据
	fmt.Printf("%d\n", a)
	fmt.Printf("%f\n", b)
	//%5d   5表示位数
	fmt.Printf("%5d\n", a)
	//不足会进行补空格  %05d  表示使用0补位
	fmt.Printf("%05d\n", a)
	//%-05d  表示使用0补位在后面补位
	fmt.Printf("%-05d\n", a)
	//如果位数大于设定的数字，则不会进行补位
	c := 1231234
	fmt.Printf("%5d\n", c)
	//%f  占位符  默认保留6位小数  %.3f表示小数点后保留三位小数(四舍五入)
	fmt.Printf("%.3f\n", b)

	//%t  占位符，表示输出一个bool类型数据
	var flag bool
	fmt.Printf("%t\n", flag)
	//%s  占位符，表示输出一个String类型数据
	var ss string = "在吗"
	fmt.Printf("%s\n", ss)
	//字符变量类型默认输出ASCII码值
	var bt byte = 'a'
	fmt.Println(bt)
	//%c 占位符，表示输出一个byte类型数据
	fmt.Printf("%c\n", bt)
}
