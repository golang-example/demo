package main

import "fmt"

func main() {
	//声明初始化一个变量
	var x int = 100
	var str string = "hello world"
	//声明初始化多个变量
	var  i, j, k int = 1, 2, 3
	//不用指明类型，通过初始化值来推导
	var b = true //bool型

	const LENGTH int = 10 //常理

	fmt.Println(x, str, i, j, k, b, LENGTH)

	abc()
}

func abc() {
	var v21 int32  //被定义初始化为0
	var v22 int = 2
	var v23 = 3 //被自动识别为int类型
	v24 := 4    //简易声明&定义的方式 等价于 var v24 int = 4;
	v21 = int32(v23) //强制转换

	g, h := 123, "hello"//注意：下行这种不带声明格式的只能在函数体中出现
	fmt.Println("v21 is", v21) //v21被赋新值
	fmt.Println("v22 is", v22)
	fmt.Println("v23 is", v23)
	fmt.Println("v24 is", v24)
	fmt.Println(g, h)
}
