package main

import "fmt"

func main() {
	demo_funcs()
	fmt.Println(add(1, 2))
	fmt.Println(add(3, 4))
	fmt.Println(swap(1, 2))
	fmt.Println(swap(3, 4))
	fmt.Println(swap2(1, 2))
	fmt.Println(swap2(3, 4))
}

func demo_funcs() {
	fmt.Println("Hello from demo_funcs")
}

// Variable type comes after the variable name
// If multiple variables have the same type, you can write it once at the end
func add(a, b int) int {
	return a + b
}

// We can return multiple values
func swap(a, b int) (int, int) {
	return b, a
}

// We can use named return values
func swap2(a, b int) (x, y int) {
	x = b
	y = a
	return
}
