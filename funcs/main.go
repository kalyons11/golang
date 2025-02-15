package main

import "fmt"

// We can use var to declare variables
var c, python, java bool

// We can pass functions as arguments
// We can return functions
// We can assign functions to variables
// We can use functions as values
// We can use functions as types
// We can use functions as structs
// We can use functions as fields in structs
// We can use functions as return values in functions
// We can use functions as arguments in functions
// We can use functions as elements in arrays
// We can use functions as elements in slices
// We can use functions as elements in maps
// We can use functions as elements in channels
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

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

	// The := syntax is shorthand for declaring and initializing a variable
	// It is only available inside functions
	k := 3
	fmt.Println(k)

	hypot := func(x, y float64) float64 {
		return x + y
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot)) // 7
	fmt.Println(compute(func(x, y float64) float64 {
		return x * y
	})) // 12
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
