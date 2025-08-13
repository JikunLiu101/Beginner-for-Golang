package main

import "fmt"

func main() {
	returnAndDefer()
}

func returnFunc() int {
	fmt.Println("returnFunc finished...")
	return 0
}

func deferFunc() int {
	fmt.Println("deferFunc finished...")
	return 0
}

func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}
