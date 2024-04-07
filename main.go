package main

import "fmt"

func plus2(n *int) {
	*n += 2
}

func main() {
	a := 8
	fmt.Println(&a)
	var n *int = &a
	fmt.Println(*n)
	fmt.Println(n)
	plus2(n)
	fmt.Println(a)
}
