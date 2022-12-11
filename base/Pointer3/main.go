package main

import "fmt"

func main() {
	n := 123
	fmt.Printf("value %d.\n", n)
	change(n)
	fmt.Printf("value %d.\n", n)
	change2(&n)
	fmt.Printf("value %d.\n", n)
}

func change(n int) {
	n += 2
}

func change2(n *int) {
	*n += 2
}
