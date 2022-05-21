package main

import "fmt"

func main() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(sum(2, 3))
	fmt.Println("exit")
}

func sum(x, y int) (sum int) {
	defer func() {
		sum *= 2
	}()

	sum = x + y
	return
}
