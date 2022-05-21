package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	go showNumbers(10)
	//runtime.Gosched()
	time.Sleep(time.Second)
	fmt.Println(runtime.NumCPU())
	fmt.Println("exit")
}

func showNumbers(num int) {
	for i := 0; i < num; i++ {
		fmt.Println(i)
	}
}
