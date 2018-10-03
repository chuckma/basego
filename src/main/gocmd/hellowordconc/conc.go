package main

import (
	"fmt"
)

func main() {
	ch :=make(chan string)// 开一个 chan
	for i:=0;i<5000 ;i++  {
		// go start goroutine 并发执行
		go printHelloWorld(i,ch)
	}

	for{
		msg := <-ch
		fmt.Println(msg)
	}

}

func printHelloWorld(i int,ch chan  string)  {
	for{
		ch <- fmt.Sprintf("hello world from goroutine %d!\n",i)

	}
}
