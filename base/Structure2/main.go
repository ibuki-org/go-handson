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

// 参照渡しじゃない方法
//func main() {
//	taro := Mydata{"Taro", []int{10, 20, 30}}
//	fmt.Println(taro)
//	taro = rev(taro)
//	fmt.Println(taro)
//}

func rev(md Mydata) Mydata {
	od := md.Data
	nd := []int{}
	for i := len(od) - 1; i >= 0; i-- {
		nd = append(nd, od[i])
	}
	md.Data = nd
	return md
}

// 参照渡しの方法
func main() {
	taro := Mydata{"Taro", []int{10, 20, 30}}
	fmt.Println(taro)
	rev2(&taro)
	fmt.Println(taro)
}

func rev2(md *Mydata) {
	od := (*md).Data
	nd := []int{}
	for i := len(od) - 1; i >= 0; i-- {
		nd = append(nd, od[i])
	}
	md.Data = nd
}
