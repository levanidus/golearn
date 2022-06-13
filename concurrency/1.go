package main

import (
	"fmt"
	"time"
)

func main() {
	go count("sheep")
	go count("fish")
	//time.Sleep(time.Second * 2)
	fmt.Scanln()
}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Microsecond * 500)
	}
}
