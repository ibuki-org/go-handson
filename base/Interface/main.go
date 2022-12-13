package main

import "fmt"

type Data interface {
	Initial(name string, data []int)
	PrintData()
}

type Mydata struct {
	Name string
	Data []int
}

func (m *Mydata) Initial(name string, data []int) {
	m.Name = name
	m.Data = data
}

func (m *Mydata) PrintData() {
	fmt.Println("Name:", m.Name)
	fmt.Println("Data:", m.Data)
}

func main() {
	var ob Data = new(Mydata) // <- interfaceからデータを作る
	//var ob Mydata = Mydata{} <- structからデータを作る
	ob.Initial("sachiko", []int{55, 66, 77})
	ob.PrintData()
}
