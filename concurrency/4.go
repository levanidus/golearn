package main

import "fmt"

func main() {
	c := make(chan string, 2)

	c <- "hello"
	c <- "yes"
	c <- "hi"

	fmt.Println(<-c)
	fmt.Println(<-c)
}
