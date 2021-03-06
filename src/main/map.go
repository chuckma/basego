package main

import "fmt"

/*
	map
	 声明变量，默认 map 是 nil
	var map_variable map[key_data_type]value_data_type

	使用 make 函数
	map_variable := make(map[key_data_type]value_data_type)
*/

func main() {
	//var countryCapitalMap map[string]string /*创建集合 */
	//countryCapitalMap = make(map[string]string)
	//
	///* map插入key - value对,各个国家对应的首都 */
	//countryCapitalMap [ "France" ] = "Paris"
	//countryCapitalMap [ "Italy" ] = "罗马"
	//countryCapitalMap [ "Japan" ] = "东京"
	//countryCapitalMap [ "India " ] = "新德里"
	//
	///*使用键输出地图值 */
	//for country := range countryCapitalMap {
	//	fmt.Println(country, "首都是", countryCapitalMap [country])
	//}
	//
	///*查看元素在集合中是否存在 */
	//captial, ok := countryCapitalMap [ "美国" ] /*如果确定是真实的,则存在,否则不存在 */
	///*fmt.Println(captial) */
	///*fmt.Println(ok) */
	//if (ok) {
	//	fmt.Println("美国的首都是", captial)
	//} else {
	//	fmt.Println("美国的首都不存在")
	//}

	deleteElement()
	atest()
}

func deleteElement()  {
	/* 创建map */
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

	fmt.Println("原始地图")

	/* 打印地图 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap [ country ])
	}

	/*删除元素*/
	delete(countryCapitalMap, "France")
	fmt.Println("法国条目被删除")

	fmt.Println("删除元素后地图")

	/*打印地图*/
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap [ country ])
	}
}

func atest()  {
	//testMap :=map[string]string{"France":"dgadsg"}
	//testMap["test"]="ggggg"
	var testMap  map[string]string
	testMap = make(map[string]string)
	testMap["test"]="ggggg"
	for key := range testMap {
		fmt.Println("testMap test :",testMap[key])

	}
}