package main

import "fmt"

func main() {

	var mapList map[string]int

	var mapAssigned map[string]int

	mapList = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapList

	// 下面两行命令打印出了相同的内存地址，说明map之间的复制是值传递
	fmt.Printf("The mapList of mapList is %p\n", mapList)
	fmt.Printf("The address of mapAssigned is %p\n", mapAssigned)

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapCreated["key3"] = 3

	fmt.Printf("Map literal as \"one\" is %d\n", mapList["one"])
	fmt.Printf("Map created as \"key2\" is %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned as \"two\" is %d\n", mapAssigned["two"])
	fmt.Printf("Map literal as \"ten\" is %d\n", mapList["ten"])

	mapList["one"] = 11
	fmt.Printf("Map literal as \"one\" is %d\n", mapList["one"])
	fmt.Printf("Map assigned as \"one\" is %d\n", mapAssigned["one"])
}
