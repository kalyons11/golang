package main

import "fmt"

func main() {
	show_arrays()
}

func show_arrays() {
	fmt.Println("Hello from show_arrays")

	var a [2]int
	a[0] = 1
	a[1] = 2
	fmt.Println(a) // [1 2]
}
