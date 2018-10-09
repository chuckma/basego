package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"]=1
	m["k2"]=2
	m["k3"]=12
	m["k4"]=0

	fmt.Println("map",m)

	v1 := m["k1"]
	fmt.Println("v1",v1)


	fmt.Println("len:", len(m))

	// map delete 不存在的key 的时候不会报错
	delete(m,"k565")

	fmt.Println("map",m)

	// 判断 map kv 是否存在
	_,prs :=m["k4"]
	fmt.Println("prs",prs)


	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
