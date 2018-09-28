package main

import "fmt"

/*
	条件语句
*/

func main() {
	/*局部变量*/

	var a int = 30

	if a < 20 {
		fmt.Printf("a 小于 20 \n")
	}else {
		/* 如果条件为 false 则执行以下语句 */
		fmt.Printf("a 不小于 20\n" );
	}
	fmt.Printf("a 的值为 : %d\n", a)
}
