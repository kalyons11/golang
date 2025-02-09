package main

import (
	"fmt"
	"math"
)

func main() {
	show_types()
	convert_types()
}

func convert_types() {
	// Go is a statically typed language, so we need to convert types explicitly
	var a int = 1
	var b float64 = float64(a)
	var c int = int(b)

	fmt.Printf("a: %T\n", a)
	fmt.Printf("b: %T\n", b)
	fmt.Printf("c: %T\n", c)

	// Go will not automatically convert types
	// var d int = 1
	// var e float64 = d // This will not compile

	// We can also convert between complex types
	var f complex64 = 1 + 1i
	var g complex128 = complex128(f)

	fmt.Printf("f: %T\n", f)
	fmt.Printf("g: %T\n", g)

	var x, y int = 3, 4
	var res float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(res)
	fmt.Println(x, y, z)

	// We can also use the short declaration operator
	h := 1
	i := float64(h)
	j := int(i)

	fmt.Printf("h: %T\n", h)
	fmt.Printf("i: %T\n", i)
	fmt.Printf("j: %T\n", j)
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

	// Constants are declared using the const keyword
	const PI = 3.14159
	const E = 2.71828
	const G = 9.81
	const truth = true

	fmt.Printf("PI: %T\n", PI)
	fmt.Printf("E: %T\n", E)
	fmt.Printf("G: %T\n", G)
	fmt.Printf("truth: %T\n", truth)

	// Numeric constants
	const (
		Big   = 1 << 100
		Small = Big >> 99
	)

	// fmt.Printf("Big: %T\n", Big) // This will not compile because of overflow
	fmt.Printf("Small: %T, %v\n", Small, Small)
	fmt.Printf("need_int(Small): %T, %v\n", need_int(Small), need_int(Small))
	fmt.Printf("need_float(Small): %T, %v\n", need_float(Small), need_float(Small))
	fmt.Printf("need_float(Big): %T, %v\n", need_float(Big), need_float(Big))
	// fmt.Printf("need_int(Big): %T\n", need_int(Big)) // This will not compile because of overflow
}

func need_int(x int) int { return x*10 + 1 }

func need_float(x float64) float64 { return x * 0.1 }
