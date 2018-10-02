package main

import "fmt"

/*
	特殊常量 iota
	iota 每当遇到 const 关键字的时候，iota 值会被重置为 0
	const 体内，每新增一行常量的时候，iota 会被计数一次
*/


const(
	/*跳值使用*/
	//b = iota  // 0
	//c = iota  // 1
	//_	// 跳值  iota +1 = 2 并赋值给 _
	//_	// 继续赋值给 _ = (2+1)
	//d = iota // 到这一行 d 就等于4 了

	/*插队使用*/
	//a = iota
	//b = 3.14
	//c = iota


	/*
		奇特的现象
		b 与 b = iota 会是不同的结果
		b 这种事表达式的隐式使用
		当没有给 b 使用表达式的时候，它会自动地向上使用第一个非空表达式
	*/
	//a = iota*2
	//b = iota	// b 与 b = iota
	//c = iota	// c 与 c = iota


	/*单行使用法*/
	a,b  = iota+1,iota+3
	c,d
	f=iota

)

func main() {

	fmt.Println("a的常量值为：", a)
	fmt.Println("b的常量值为：", b)
	fmt.Println("c的常量值为：", c)
	fmt.Println("d的常量值为：", d)
	fmt.Println("f的常量值为：", f)

}

