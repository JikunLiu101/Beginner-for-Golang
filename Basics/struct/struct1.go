package main

import "fmt"

// declare a new data type
type myInt int

// declare a struct
type Book struct {
	title  string
	author string
	pages  myInt
}

func main() {
	var a myInt = 10
	fmt.Println("a = ", a)
	fmt.Printf("type of a is %T\n", a)

	var book1 Book
	book1.title = "Go Programming Language"
	book1.author = "www.golang.com"
	book1.pages = 20

	fmt.Printf("%v\n", book1)

	changeBook1(book1)
	fmt.Printf("%v\n", book1)

	changeBook2(&book1)
	fmt.Printf("%v\n", book1)
}

func changeBook1(book Book) {
	book.author = "www.google.com"
}

func changeBook2(book *Book) {
	book.author = "www.google.com"
}
