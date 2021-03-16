package main

import "fmt"

func main() {
	//按照顺序对应定义变量赋值
	//多重定义:  变量1, 变量2...  :=  值1,值2......
	a, b, c, d := 1, 2, 3, 4
	fmt.Println(a, b, c, d)

	//交换变量的值
	//方法一: 中间变量交换
	temp := a
	a = b
	b = temp
	fmt.Println(a, b)
	//方法二: 加减运算交换
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)
	//方法三:多重赋值交换
	a, b = b, a
	fmt.Println(a, b)

	//匿名变量  使用 _ 来声明一个匿名变量，多用于接受函数返回的不想要的变量
	f, g, _, z := 5, 6, 7, 8
	fmt.Println(f, g, z)
}
