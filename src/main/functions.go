package main

/*
	函数


	func function_name( [parameter list] ) [return_types] {
	   函数体
	}

	func：函数由 func 开始声明
	function_name：函数名称，函数名和参数列表一起构成了函数签名。
	parameter list：参数列表，参数就像一个占位符，当函数被调用时，
		你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，
		也就是说函数也可以不包含参数。
	return_types：返回类型，函数返回一列值。return_types 是该列值的数据类型。
		有些功能不需要返回值，这种情况下 return_types 不是必须的。
		函数体：函数定义的代码集合。
*/

func main() {

	println(max(3, 4))

	a, b := swap("123", "321")
	println(a,b)
}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 声明局部变量 */
	var result int

	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}

/* Go 函数可以返回多个值,例如：*/
func swap(x, y string) (string, string) {
	return y, x
}
