package main

import "fmt"

func main() {
	a := 5
	b := 10
	fmt.Println("Before swap:", a, b)
	swap(a, b)
	fmt.Println("After swap:", a, b)
	swapPointer(&a, &b)
	fmt.Println("After swap pointer:", a, b)
}

func swapPointer(a *int, b *int) {
	var temp int

	temp = *a
	*a = *b
	*b = temp
}

func swap(a int, b int) {
	var temp int

	temp = a
	a = b
	b = temp
}
