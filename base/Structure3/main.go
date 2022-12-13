package main

import "fmt"

type Mydata struct {
	Name string
	Data []int
}

func main() {
	taro := new(Mydata)
	fmt.Println(taro)
	taro.Name = "Taro"
	taro.Data = make([]int, 10, 20)
	fmt.Println(taro)
}
