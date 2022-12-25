package main

import (
	"fmt"
	"time"
)

func total(c chan int) {
	n := <-c
	fmt.Println("n = ", n)
	t := 0
	for i := 0; i <= n; i++ {
		t += i
	}
	fmt.Println("total:", t)
}

func main() {
	c := make(chan int)
	// ここでやると、deadlockになる
	// c <- 100
	go total(c)
	// ここでやると、正常終了する
	c <- 100
	time.Sleep(100 * time.Millisecond)
}
