package main

import "fmt"

func total(n int, c chan int) {
	t := 0
	for i := 0; i < n; i++ {
		t += i
	}
	c <- t
}

func main() {
	c := make(chan int)
	go total(1000, c)
	go total(100, c)
	go total(10, c)
	// 順番に表示されるわけではなく、チャンネル内で実行が終わった順に表示される
	x, y, z := <-c, <-c, <-c
	fmt.Println(x, y, z)
}
