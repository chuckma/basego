package main

import "fmt"

func test() {
}
func main() {

	//注释
	// fmt.Println("Hello, World!")
	/*变量申明方式*/
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)
	/*多变量申明*/
	// g, h := 123, "hello"
	//t := "dsafds"
	//t1 := false
	// println(x, y, a, b, c, d, e, f, g, h, t)
	//println(t)
	//println(t1)
	//println(C3, C4, C5)

	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)

}

//func init() {
// fmt.Println("首先被执行！")
//}

/* 变量的申明方式*/
//var a = "根据值自动判断变量类型的申明"
//var b string = "带类型的变量申明"
//var c = false

/*多变量申明*/

var x, y int // int 默认是 0
var ( // 这种因式分解关键字的写法一般用于声明全局变量
	a int
	b bool // 默认是 false
)

var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

/*
常量,程序运行中不被修改的量
数据类型可以是布尔，数字，字符串

*/
// 显式定义
const C1 string = "显式定义常量"

// 隐式定义
const C2 = "隐式定义常量"

// 多个相同类型常量 简写
const C3, C4, C5 = 1, 2, 3
