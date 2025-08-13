package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	// reference copy
	s1 := s[0:1]
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(s1), cap(s1), s1)

	s1[0] = 100
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(s1), cap(s1), s1) // [100, 2]
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(s), cap(s), s)    // [100, 2, 3]

	// deep copy
	s2 := make([]int, 3) // [0, 0, 0]
	copy(s2, s)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(s2), cap(s2), s2) // [100, 2, 3]
}
