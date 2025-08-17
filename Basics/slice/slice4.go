package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)

	stus := []student{
		{name: "Bob", age: 30},
		{name: "Alice", age: 23},
		{name: "Jack", age: 21},
	}

	for _, stu := range stus {
		m[stu.name] = &stu //stu is one variable, so the address will always be the same
	}

	fmt.Println(m)

	for i := 0; i < len(stus); i++ {
		m[stus[i].name] = &stus[i]
	}

	fmt.Println(m)

	for _, stu := range stus { //这个遍历方法其实和java的for each一样是创建了一个副本
		stu.age = stu.age + 1
	}

	fmt.Println(stus)

	for k, v := range m {
		println(k, "->", v.name)
	}
}
