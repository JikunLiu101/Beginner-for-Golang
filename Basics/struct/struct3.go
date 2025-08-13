package main

import "fmt"

type Human struct {
	name   string
	gender string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

type SuperMan struct {
	Human // Superman inherits Human
	level int
}

func (this *SuperMan) Eat() { // SuperMan inherits Eat() from Human
	fmt.Println("SuperMan.Eat()...")
}

func (this *SuperMan) Fly() { // SuperMan has new child menthod
	fmt.Println("SuperMan.Fly()...")
}

func main() {
	h := new(Human)
	h.name = "Jay"
	h.gender = "f"
	h.Walk()
	h.Eat()

	fmt.Println("===========")
	s := SuperMan{Human{"Leo", "m"}, 5}
	s.Eat()
	s.Fly()
	s.Walk() // child can call parent method

	fmt.Println("===========")
	var s2 SuperMan
	s2.name = "Jeff"
	s2.gender = "m"
	s2.level = 3
	s2.Fly()
}
