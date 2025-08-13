package main

import "fmt"

type Hero struct { // "H" is capitalised, means this struct is public
	Name  string
	_Ad   int // this _Ad attribute is public, underscore has no special effects.
	level int // this level attribute is private, not public
}

// (this Hero) means this function is tied to Hero struct
func (this Hero) GetName() string {
	return this.Name
}

// use * to make sure reference copy, otherwise it will be value copy
func (this *Hero) SetName(name string) {
	this.Name = name
}

func (this Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("_Ad = ", this._Ad)
	fmt.Println("Level = ", this.level)
}

func main() {
	hero := Hero{"Leo", 100, 1}

	hero.Show()
	hero.SetName("Jay")
	hero.Show()

}
