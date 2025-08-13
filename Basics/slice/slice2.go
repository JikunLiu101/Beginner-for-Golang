package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// add element 1 to slice numbers, now len is 4, cap is still 5
	numbers = append(numbers, 1)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// add element 2 to slice numbers, now len is 5, cap is 5
	numbers = append(numbers, 2)

	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	// add element 3 to slice numbers, now len is 6, cap is 10 (cap will double)
	numbers = append(numbers, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 5)
	numbers = append(numbers, 6)
	numbers = append(numbers, 7)
	numbers = append(numbers, 8)
	numbers = append(numbers, 9)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	var numbers2 = make([]int, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)

	numbers2 = append(numbers2, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers2), cap(numbers2), numbers2)
}
