package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("一个空的切片",s)
	s[0]="a"
	s[1]="b"
	s[2]="c"
	//s[3]="d"
	fmt.Println("s",s)

	// 添加一个元素
	s = append(s, "d")
	// 添加多个元素
	s = append(s, "e", "f")
	fmt.Println("s",s)


	// slice 的 copy 创建一个长度和 s 一样的空 slice
	c := make([] string, len(s))
	copy(c, s)
	fmt.Println("从 s copy 来的 ",c)

	// 切片的切割操作,左包含，右不包含

	l := s[2:5]
	fmt.Println("l",l)
	// 去掉切片最后一个元素
	l = s[:5]
	fmt.Println("l",l)

	l = s[2:]
	fmt.Println("sl3:", l)




}
