package main

import "fmt"

// 本质是一个指针
type Animal interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	color string
}

// 必须重写interface里的所有方法才算实现了这个interface

func (this *Cat) Sleep() { //由于interface本质上是指针，所以这里是 (this *Cat)
	fmt.Println("Cat sleeps....")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

// Dog也实现了这个接口

func (this *Dog) Sleep() {
	fmt.Println("Dog sleeps....")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal Animal) {
	animal.Sleep() //多态
	fmt.Println("color: ", animal.GetColor())
	fmt.Println("type: ", animal.GetType())
}

func main() {
	var animal Animal // animal 是一个接口
	animal = &Cat{color: "red"}
	animal.Sleep() // 调用Cat的Sleep()

	fmt.Println("===========")
	animal = &Dog{color: "red"}
	animal.Sleep() // 调用Dog的Sleep()

	fmt.Println("===========")
	cat := Cat{color: "red"}
	fmt.Println("===========")
	dog := Dog{color: "yellow"}
	showAnimal(&cat) //由于Animal是interface，是指针，所以这里要传地址
	showAnimal(&dog)
}
