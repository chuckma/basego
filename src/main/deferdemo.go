package main

import "fmt"

/*
	defer

	https://studygolang.com/articles/10167


	在golang当中，defer代码块会在函数调用链表中增加一个函数调用。这个函数调用不是普通的函数调用，
	而是会在函数正常返回，也就是return之后添加一个函数调用。因此，defer通常用来释放函数内部变量。



*/

func main() {
	for i:=0 ;i<5;i++{
		defer fmt.Printf("%d",i)
		fmt.Println("bbbbb")
	}
	fmt.Println("aaaaa")
}
