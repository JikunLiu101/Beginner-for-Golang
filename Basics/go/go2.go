package main

import "fmt"

type People struct {
}

func (people *People) ShowA() {
	fmt.Println("ShowA")
	people.ShowB()
}

func (people *People) ShowB() {
	fmt.Println("ShowB")
}

type Teacher struct {
	People
}

func (teacher *Teacher) ShowB() {
	fmt.Println("TeacherShowB")
}

func main() {
	t := Teacher{}
	t.ShowA() // showA showB
}
