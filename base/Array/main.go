package main

import (
	"fmt"
	"go-handson/hello"
	"strconv"
	"strings"
)

func main() {
	x := hello.Input("input data")
	ar := strings.Split(x, ",")
	t := 0
	for i := 0; i < len(ar); i++ {
		n, err := strconv.Atoi(ar[i])
		if err != nil {
			goto err
		}
		t += n
	}
	fmt.Println("total: ", t)
	return
err:
	fmt.Println("error")
}
