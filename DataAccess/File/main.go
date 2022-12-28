package main

import (
	"fmt"
	"go-handson/hello"
	"os"
)

func main() {
	wt := func(f *os.File, s string) {
		_, er := f.WriteString(s + "\n")
		if er != nil {
			fmt.Println(er)
			return
		}
	}

	fn := "data.txt"

	f, er := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if er != nil {
		fmt.Println(er)
		f.Close()
		return
	}
	fmt.Println("*** start ***")
	wt(f, "*** start ***")
	for {
		s := hello.Input("type message")
		if s == "" {
			break
		}
		wt(f, s)
	}
	wt(f, "*** end ***\n\n")
	fmt.Println("*** end ***")
	er = f.Close()
	if er != nil {
		fmt.Println(er)
	}
}
