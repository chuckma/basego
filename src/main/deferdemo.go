package main

import (
	"fmt"
	"os"
)

/*
	defer

	https://studygolang.com/articles/10167


	在golang当中，defer代码块会在函数调用链表中增加一个函数调用。这个函数调用不是普通的函数调用，
	而是会在函数正常返回，也就是return之后添加一个函数调用。因此，defer通常用来释放函数内部变量。


	Defer 被用来确保一个函数调用在程序执行结束前执行。
	同样用来执行一些清理工作。 defer 用在像其他语言中的ensure 和 finally用到的地方

*/

func main() {


	// 在 closeFile 后得到一个文件对象，我们使用 defer通过 closeFile 来关闭这个文件。
	// 这会在封闭函数（main）结束时执行，就是 writeFile 结束后。

	f := createFile("/Users/lucasma/Desktop/defer.txt")
	writeFile(f)
	defer closeFile(f)
}
func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}
func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data213123")
}
func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
