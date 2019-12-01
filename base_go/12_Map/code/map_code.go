//package main
//
//import "fmt"
//
//func main() {
//	var countryCapitalMap map[string]string /*创建集合 */
//	// 此时map为nil,不能再下面进行赋值
//	fmt.Println(countryCapitalMap)
//	// 必须使用make()进行初始化才能进行赋值
//	countryCapitalMap = make(map[string]string)
//
//	/* map插入key - value对,各个国家对应的首都 */
//	countryCapitalMap [ "France" ] = "巴黎"
//	countryCapitalMap [ "Italy" ] = "罗马"
//	countryCapitalMap [ "Japan" ] = "东京"
//	countryCapitalMap [ "India " ] = "新德里"
//
//	/*使用键输出地图值 */
//	for country := range countryCapitalMap {
//		// 当range一个map时,只有一个值就是key,而不像数组为一个索引,一个值
//		fmt.Println(country)
//		fmt.Println(country, "首都是", countryCapitalMap [country])
//	}
//
//	/*查看元素在集合中是否存在 */
//	fmt.Println(countryCapitalMap)
//	// 指定要查找的country
//	var country string = "Japan"
//	capital, ifTrue:= countryCapitalMap[country] /*如果确定是真实的,则存在,否则不存在 */
//	// 当要取的map对应的key不存在时,返回"", 若map对应的key存在,正好也是空值,那如何判断key到底存不存在呢?
//	// 可以根据如上的ifTrue来进行判断,ifTrue是一个bool值
//	fmt.Println(capital)
//	fmt.Println(ifTrue)
//	// 可以根据if判断bool值来进行判断key是否存在
//	if ifTrue {
//		fmt.Printf("%s 的首都是 %s", country, capital)
//	} else {
//		fmt.Printf("%s的首都不存在", country)
//	}
//}


package main

import "fmt"

func main() {
	/* 创建map */
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
	fmt.Println("原始地图")

	// 国家
	for country := range countryCapitalMap {
		fmt.Println(country)
	}
	/* 打印地图 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	/*删除元素*/
	delete(countryCapitalMap, "France")
	fmt.Println("法国条目被删除")

	fmt.Println("删除元素后地图")

	/*打印地图*/
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
}