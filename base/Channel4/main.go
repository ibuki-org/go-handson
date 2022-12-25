package main

import (
	"fmt"
	"strconv"
	"time"
)

func prmsg(n int, s string) {
	// firstで変更されたchannelの値を出力している
	fmt.Println(s)
	time.Sleep(time.Duration(n) * time.Second)
}

func first(n int, c chan string) {
	const nm string = "first-"
	for i := 0; i < 10; i++ {
		s := nm + strconv.Itoa(i)
		go prmsg(n, s)
		c <- s
	}
}

func second(n int, c chan string) {
	for i := 0; i < 10; i++ {
		go prmsg(n, "second: ["+<-c+"]")
	}
}

func main() {
	c := make(chan string)
	go first(10, c)
	second(10, c)
	fmt.Println()
}
