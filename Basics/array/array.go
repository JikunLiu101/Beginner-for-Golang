package main

import "fmt"

func printArray(myArray [5]int) {
	// value copy
	for index, value := range myArray {
		fmt.Println("index = ", index, "value = ", value)
	}
}

func printSlice(mySlice []int) {
	// reference copy
	for _, value := range mySlice { // _ is blank identifier, can assign something to _ to ignore it
		fmt.Println("value = ", value)
	}
	mySlice[0] = 100
}

func main() {
	var myArray1 [10]int
	myArray2 := [5]int{1, 2, 3, 4}
	myArray3 := [4]int{11, 22, 33, 44}
	myArray := []int{1, 2, 3} // slice

	//for i := 0; i < len(myArray1); i++ {
	//	fmt.Println(myArray1[i])
	//}

	//for index, value := range myArray1 {
	//	fmt.Println("index", index, "value", value)
	//}

	fmt.Printf("myArray1: types = %T\n", myArray1)
	fmt.Printf("myArray2: types = %T\n", myArray2)
	fmt.Printf("myArray3: types = %T\n", myArray3)
	fmt.Printf("myArray: types = %T\n", myArray)

	printArray(myArray2)
	//printArray(myArray1) // cannot use myArray1 (variable of type [10]int) as [5]int value in argument to printArray
	//printArray(myArray3) // cannot use myArray3 (variable of type [4]int) as [5]int value in argument to printArray

	printSlice(myArray)
	fmt.Println("====")
	printSlice(myArray)
}
