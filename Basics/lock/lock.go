package main

import (
	"fmt"
	"strconv"
	"sync"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	ua.Lock()
	defer ua.Unlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func main() {
	var wg sync.WaitGroup
	wg.Add(30)
	ua := UserAges{ages: make(map[string]int), Mutex: sync.Mutex{}}

	for i := 0; i < 5; i++ {
		go func() {
			ua.Add(strconv.Itoa(i), i)
		}()
		for j := 0; j < 5; j++ {
			go func() {
				age := ua.Get(strconv.Itoa(j))
				fmt.Println(age)
			}()
		}
	}
}
