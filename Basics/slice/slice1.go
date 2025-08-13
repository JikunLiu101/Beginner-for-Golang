package main

import (
	"fmt"
)

func main() {
	// declare a slice and initialise
	slice1 := []int{1, 2, 3}

	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)

	// declare a slice but no space no value
	var slice2 []int

	fmt.Printf("len = %d, slice = %v\n", len(slice2), slice2)

	//slice2[0] = 1 // runtime error: index out of range [0] with length 0
	slice2 = make([]int, 3)
	slice2[0] = 10

	fmt.Printf("len = %d, slice = %v\n", len(slice2), slice2)

	// declare a slice and initilise size to 3
	slice3 := make([]int, 5)
	fmt.Printf("len = %d, slice = %v\n", len(slice3), slice3)

	var slice4 []int
	if slice4 == nil {
		fmt.Println("slice4 is nil")
	} else {
		fmt.Println("slice4 is not nil")
	}
}
