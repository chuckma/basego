package main

import "fmt"

/*
	指针
	我们都知道，变量是一种使用方便的占位符，用于引用计算机内存地址。

	Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址

	什么是指针
	一个指针变量指向了一个值的内存地址。

	类似于变量和常量，在使用指针前你需要声明指针。指针声明格式如下：
	var var_name *var-type
	var-type 为指针类型，var_name 为指针变量名，* 号用于指定变量是作为一个指针。以下是有效的指针声明
	// 指向整型
	var ip *int
	// 指向浮点型
	var fp *float32


	指针使用流程：

	定义指针变量。
	为指针变量赋值。
	访问指针变量中指向地址的值。
	在指针类型前面加上 * 号（前缀）来获取指针所指向的内容。


	Go 空指针
	当一个指针被定义后没有分配到任何变量时，它的值为 nil。

	nil 指针也称为空指针。

	nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。

	一个指针变量通常缩写为 ptr。


	指针数组

	指针作为函数参数
	Go 语言允许向函数传递指针，只需要在函数定义的参数上设置为指针类型即可。



	指针指向的指针
	如果一个指针变量存放的又是另一个指针变量的地址，则称这个指针变量为指向指针的指针变量。

	当定义一个指向指针的指针变量时，第一个指针存放第二个指针的地址，第二个指针存放变量的地址：
	var ptr **int;
*/

const MAX int = 3

func main()  {
	var a int = 10 /*声明实际变量*/
	var ip *int /*声明指针变量*/
	ip = &a /*指针变量的内存地址*/
	fmt.Printf("变量的地址: %x\n",&a)

	fmt.Printf("ip 变量存储的指针地址是：%x\n",ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip )

	// 空指针

	var  ptr *int

	fmt.Printf("ptr 的值为 : %x\n", ptr  )

	// 空指针的判断
	if ptr != nil{
		fmt.Printf("ptr 不是空指针")
	}
	if ptr == nil{
		fmt.Printf("ptr 是空指针\n")
	}



	/*指针数组 */
	b := []int{10,100,200}
	var i int
	var ptr1 [MAX]*int;

	// 整数地址赋值给指针数组
	for i=0;i<MAX ;i++ {
		ptr1[i] = &b[i]
	}

	for  i = 0; i < MAX; i++ {
		fmt.Printf("b[%d] = %d\n", i,*ptr1[i] )
	}

	/*指针作为函数参数*/
	pointerAsParam()
	/*指针指向指针*/
	pointerToPointer()
}


func pointerAsParam() {
	fmt.Println("==================")

	/* 定义局部变量 */
	var a int = 100
	var b int= 200

	fmt.Printf("交换前 a 的值 : %d\n", a )
	fmt.Printf("交换前 b 的值 : %d\n", b )

	/* 调用函数用于交换值
	* &a 指向 a 变量的地址
	* &b 指向 b 变量的地址
	*/
	swap(&a, &b);

	fmt.Printf("交换后 a 的值 : %d\n", a )
	fmt.Printf("交换后 b 的值 : %d\n", b )
}

func swap(x *int, y *int) {
	var temp int
	temp = *x    /* 保存 x 地址的值 */
	*x = *y      /* 将 y 赋值给 x */
	*y = temp    /* 将 temp 赋值给 y */
}

func pointerToPointer()  {
	fmt.Println("==================")
	var a int
	var ptr *int
	var pptr **int

	a = 3000

	/* 指针 ptr 地址 */
	ptr = &a

	/* 指向指针 ptr 地址 */
	pptr = &ptr

	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d\n", a )
	fmt.Printf("指针变量 *ptr = %d\n", *ptr )
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
}
/*
	Go 空指针
	当一个指针被定义后没有分配到任何变量时，它的值为 nil。

	nil 指针也称为空指针。

	nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。

	一个指针变量通常缩写为 ptr。

	查看以下实例：
*/

