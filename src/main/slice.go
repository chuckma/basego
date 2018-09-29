package main

import "fmt"

/*
	切片
	定义切片 切片不需要说明长度。
	var identifier []type

*/

func main() {
	//var numbers = make([]int,3,5)
	//
	//printSlice(numbers)
	//
	//// 空切片
	//var numbers1 []int
	//
	//nilSlice(numbers1)
	//
	//if(numbers1 == nil){
	//	fmt.Printf("切片是空的")
	//}
	//创建切片
	numbers :=[] int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	splitSlice(numbers)

	// 切片添加元素
	addElement()
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
/*空(nil)切片*/
func nilSlice(x []int)  {
	var numbers []int

	printSlice(numbers)

	if(numbers == nil){
		fmt.Printf("切片是空的")
	}
}

/*切片截取*/

func splitSlice(numbers []int)  {



	printSlice(numbers)

	// 输出原始切片
	fmt.Println("numbers :" , numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4]==",numbers[1:4])
	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])
	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])


	fmt.Println("====================")
	numbers1 := make([]int,0,5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)


}

/*append() 和 copy() 函数*/

func addElement()  {
	fmt.Println("=========切片添加元素===========")
	var num []int
	printSlice(num)

	/* 允许追加空切片 */
	num = append(num, 0)
	printSlice(num)

	/* 向切片添加一个元素 */
	num = append(num, 1)
	printSlice(num)

	/* 同时添加多个元素 */
	num = append(num, 2,3,4)
	printSlice(num)

	/* 创建切片 num1 是之前切片的两倍容量*/
	num1 := make([]int, len(num), (cap(num))*2)

	/* 拷贝 num 的内容到 num1 */
	copy(num1,num)
	printSlice(num1)
}