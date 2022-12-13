package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Data interface {
	SetValue(value map[string]string)
	PrintData()
}

type Mydata struct {
	Name string
	Data []int
}

func (md *Mydata) SetValue(vals map[string]string) {
	md.Name = vals["name"]
	valt := strings.Split(vals["data"], " ")
	vali := []int{}
	for _, v := range valt {
		n, _ := strconv.Atoi(v)
		vali = append(vali, n)
	}
	md.Data = vali
}

func (md *Mydata) PrintData() {
	fmt.Println("Name: ", md.Name)
	fmt.Println("Data: ", md.Data)
}

type Yourdata struct {
	Name string
	Mail string
	Age  int
}

func (md *Yourdata) SetValue(vals map[string]string) {
	md.Name = vals["name"]
	md.Mail = vals["mail"]
	md.Age, _ = strconv.Atoi(vals["age"])
}

func (md *Yourdata) PrintData() {
	fmt.Println("Name: ", md.Name)
	fmt.Println("Mail: ", md.Mail)
	fmt.Println("Age: ", md.Age)
}

func main() {
	ob := [2]Data{}
	ob[0] = new(Mydata)
	ob[0].SetValue(map[string]string{"name": "Taro", "data": "1 2 3 4 5"})
	ob[1] = new(Yourdata)
	ob[1].SetValue(map[string]string{"name": "Jiro", "mail": "yugo.ibuki@gmail.com", "age": "20"})
	for _, v := range ob {
		v.PrintData()
		fmt.Println()
	}
}
