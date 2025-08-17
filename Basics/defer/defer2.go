package main

func main() {
	defer_call()
}

func defer_call() {
	defer println("call 1")
	defer println("call 2")
	defer println("call 3")
	panic("panic")
}
