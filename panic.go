package main

import "fmt"

func main() {
	makePanic()
}

func makePanic() {
	defer func() {
		panicValue := recover()
		fmt.Println(panicValue)
	}()

	panic("some panic")
	fmt.Println("unreachable code")
}
