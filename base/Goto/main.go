package main

import (
	"fmt"
	"go-handson/hello"
	"strconv"
)

func main() {
	t := 0
	x := hello.Input("type a number")
	n, err := strconv.Atoi(x)
	if err != nil {
		/* err というラベルに対して処理を飛ばす事ができる */
		goto err
	}
	for i := 1; i <= n; i++ {
		t += i
	}
	fmt.Println("total: ", t)
	return
err:
	fmt.Println("error")
}
