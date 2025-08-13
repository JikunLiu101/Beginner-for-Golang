package main

import "fmt"

func main() {
	cityMap := make(map[string]string)

	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["India"] = "New Delhi"

	// interation

	for key, value := range cityMap {
		fmt.Println("key:", key, "value:", value)
	}

	// delete
	delete(cityMap, "Japan")

	// update
	cityMap["India"] = "New New Delhi"

	fmt.Println("==========")
	printMap(cityMap)

	fmt.Println("==========")
	cityMap["England"] = "Loooooondon"
	changeValue(cityMap)
	printMap(cityMap)
}

func printMap(m map[string]string) {
	// reference copy
	for key, value := range m {
		fmt.Println("key:", key, "value:", value)
	}
}

func changeValue(m map[string]string) {
	// reference copy
	m["England"] = "London"
}
