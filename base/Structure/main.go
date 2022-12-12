package main

import "fmt"

//var mydata struct {
//	Name string
//	Data []int
//}

type Mydata struct {
	Name string
	Data []int
}

func main() {
	taro := Mydata{"Taro", []int{10, 20, 30}}
	hanako := Mydata{"Hanako", []int{40, 50, 60}}
	fmt.Println(taro)
	fmt.Println(hanako)
}
