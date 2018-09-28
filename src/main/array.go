package main

import "fmt"

/*
	数组
	申明
	var variable_name [SIZE] variable_type
	eg. 定义了数组 balance 长度为 10 类型为 float32：
		var balance [10] float32

	init array eg.
	var balance = [5]float32{1.12,2,232,2,12}
	初始化数组中 {} 中的元素个数不能大于 [] 中的数字。
	如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小：
*/

//以下演示了数组完整操作（声明、赋值、访问）的实例：

func main() {

	var n [10]int
	var i, j int

	// 为 数组 n 初始化 元素
	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}


	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j] )
	}

}
