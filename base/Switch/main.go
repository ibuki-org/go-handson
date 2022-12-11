package main

import (
	"fmt"
	"go-handson/hello"
	"strconv"
)

func main() {
	x := hello.Input("type a number")
	fmt.Print(x + "月は、")
	/* strconv.Atoiのerr処理は以下のように省略できる */
	switch n, err := strconv.Atoi(x); n {
	case 0:
		fmt.Println("整数値がえられません")
		fmt.Println(err)
	case 1, 2, 12:
		fmt.Println("冬")
	case 3, 4, 5:
		fmt.Println("春")
	case 6, 7, 8:
		fmt.Println("夏")
	case 9, 10, 11:
		fmt.Println("秋")
	default:
		fmt.Println("月の値ではありません。")
	}
}
