package main

import "fmt"

// We can use var to declare variables
var c, python, java bool

func main() {
	demo_funcs()
	fmt.Println(add(1, 2))
	fmt.Println(add(3, 4))
	fmt.Println(swap(1, 2))
	fmt.Println(swap(3, 4))
	fmt.Println(swap2(1, 2))
	fmt.Println(swap2(3, 4))
	var x int
	_ = x
	var i, j int = 1, 2
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
	var name = "John"
	fmt.Println(name)
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
