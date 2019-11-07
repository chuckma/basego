package main

import "fmt"

/*
	结构体是由一系列相同数据类型或不同数据类型构成的数据结合，它表示一项纪录

	结构体的定义需要 type 和 struct 语句，struct 定义一个新的数据类型，结构体中
	有一个或多个成员，type 定义结构体的名称

	type struct_variable_type struct {
	   member definition;
	   member definition;
	   ...
	   member definition;
	}
*/

type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}

func main() {
	test()
	//test1()
	//test2()
	//test3()
}
func test() {
	// 创建一个新的结构体
	fmt.Println(Books{"Go语言", "lucas ma", "Go语言基础", 123})

	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "Go语言", author: "lucas ma", subject: "Go语言基础", bookId: 124})

	// 忽略的字段为 0 或 空
	fmt.Println(Books{title: "Go语言", author: "lucas ma"})

	books := Books{"Go语言", "lucas ma", "Go语言基础", 123123}
	fmt.Println(books.author)
}
func test1() {
	var Book1 Books /* 声明 Book1 为 Books 类型 */
	var Book2 Books /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "lucas ma"
	Book1.subject = "Go 语言教程"
	Book1.bookId = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "lucas ma"
	Book2.subject = "Python 语言教程"
	Book2.bookId = 6495700

	/* 打印 Book1 信息 */
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 bookId : %d\n", Book1.bookId)

	/* 打印 Book2 信息 */
	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 bookId : %d\n", Book2.bookId)
}

/*结构体作为函数参数*/
func test2() {
	var Book1 Books /* 声明 Book1 为 Books 类型 */
	var Book2 Books /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "lucas ma"
	Book1.subject = "Go 语言教程"
	Book1.bookId = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "lucas ma"
	Book2.subject = "Python 语言教程"
	Book2.bookId = 6495700

	/* 打印 Book1 信息 */
	printBook(Book1)

	/* 打印 Book2 信息 */
	printBook(Book2)

}

func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title);
	fmt.Printf("Book author : %s\n", book.author);
	fmt.Printf("Book subject : %s\n", book.subject);
	fmt.Printf("Book bookId : %d\n", book.bookId);
}

/*
	结构体指针
	var struct_pointer *Books

*/
func test3()  {
	var Book1 Books        /* Declare Book1 of type Book */
	var Book2 Books        /* Declare Book2 of type Book */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "lucas ma"
	Book1.subject = "Go 语言教程"
	Book1.bookId = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "lucas ma"
	Book2.subject = "Python 语言教程"
	Book2.bookId = 6495700

	/* 打印 Book1 信息 */
	printBookByPointer(&Book1)

	/* 打印 Book2 信息 */
	printBookByPointer(&Book2)
}

func printBookByPointer( book *Books) {
	fmt.Printf( "Book title : %s\n", book.title);
	fmt.Printf( "Book author : %s\n", book.author);
	fmt.Printf( "Book subject : %s\n", book.subject);
	fmt.Printf( "Book bookId : %d\n", book.bookId);
}