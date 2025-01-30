package main

import "fmt"

func main() {
	show_types()
}

func show_types() {
	bool_var := true
	int_var := 1
	string_var := "Hello, World!"
	float_var := 1.0
	complex_var := 1 + 1i

	// %T is used to print the type of the variable
	fmt.Printf("bool_var: %T\n", bool_var)
	fmt.Printf("int_var: %T\n", int_var)
	fmt.Printf("string_var: %T\n", string_var)
	fmt.Printf("float_var: %T\n", float_var)
	fmt.Printf("complex_var: %T\n", complex_var)

	// We can use a factored block to declare multiple variables
	var (
		a int    = 1
		b string = "Hello, World!"
		c bool   = true

		d, e, f int = 1, 2, 3
	)

	fmt.Printf("a: %T\n", a)
	fmt.Printf("b: %T\n", b)
	fmt.Printf("c: %T\n", c)
	fmt.Printf("d: %T\n", d)
	fmt.Printf("e: %T\n", e)
	fmt.Printf("f: %T\n", f)

	// Uninitialized variables are set to their zero value
	var g int        // 0
	var h string     // ""
	var i bool       // false
	var j float64    // 0.0
	var k complex128 // 0 + 0i

	fmt.Printf("g: %T\n", g)
	fmt.Printf("h: %T\n", h)
	fmt.Printf("i: %T\n", i)
	fmt.Printf("j: %T\n", j)
	fmt.Printf("k: %T\n", k)
}
