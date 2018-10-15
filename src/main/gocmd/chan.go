package main

import "fmt"

/*通道*/

func main() {

	message := make(chan string)

	go func() { message <- "ping" }()

	//msg := <-message
	fmt.Println(<-message)
}
