package main

import "fmt"

func main() {
	demo_funcs()
	fmt.Println(add(1, 2))
	fmt.Println(add(3, 4))
}

func demo_funcs() {
	fmt.Println("Hello from demo_funcs")
}

// Variable type comes after the variable name
// If multiple variables have the same type, you can write it once at the end
func add(a, b int) int {
	return a + b
}
