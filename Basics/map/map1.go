package main

import "fmt"

func main() {
	// declare type for a map, key is string, value is int
	var map1 map[string]int
	if map1 == nil {
		fmt.Println("map1 == nil")
	}

	// declare space for map
	map1 = make(map[string]int, 10)
	map1["Java"] = 10
	map1["C++"] = 20
	map1["Python"] = 30
	fmt.Println(map1) // result is unordered, because allocation is hashed

	// another declaration
	map2 := make(map[string]int)
	map2["Java"] = 100
	map2["C++"] = 200
	map2["Python"] = 300
	fmt.Println(map2)

	// third declaration
	map3 := map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Println(map3)

}
