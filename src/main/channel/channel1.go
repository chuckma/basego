package main

import (
	"fmt"
	"reflect"
)

func main() {
	nb := make(chan string, 10)
	nb <- "dsafgasfg213131"
	value := <-nb
	fmt.Println(value)
	fmt.Println(reflect.TypeOf(value))
}
