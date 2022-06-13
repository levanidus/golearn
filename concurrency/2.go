package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		count("sheep")
		wg.Done()
	}()
	go func() {
		count("fish")
		wg.Done()
	}()
	wg.Wait()
}

func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Microsecond * 500)
	}
}
