package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan string)
	go count("sheep", c)

	for {
		msg, open := <-c
		if !open {
			break
		}
		fmt.Println(msg)
	}

}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Microsecond * 500)
	}

	close(c)
}
