package main

import (
	"fmt"
	"go-handson/hello"
	"strconv"
)

func main() {
	x := hello.Input("type a price")
	n, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("error")
		return
	}
	p := float32(n)
	fmt.Println(int(p * 1.1))
}
