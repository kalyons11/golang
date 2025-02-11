package main

import (
	"fmt"
	"math"
)

func main() {
	show_for()
}

func show_for() {
	fmt.Println("Hello from show_for")

	// For loops in golang have 3 parts: init, condition, post
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// We can create a doubling loop
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// Let's call our sqrt function
	fmt.Println(sqrt(2), sqrt(-4)) // 1.4142135623730951 2i

	// We can use a short statement to declare variables in an if
	// This variable is only available in the if block
	if v := 11.0; v < 10 {
		fmt.Println(v)
	} else {
		fmt.Println(v, "is greater than 10") // We can use v here
	}

	// We can't use v here
	// fmt.Println(v)

	// Let's call our newton sqrt function and compare it to the math.Sqrt function
	fmt.Println(sqrt_newton(2), math.Sqrt(2))
	// Compute the difference between the two
	fmt.Println(math.Abs(sqrt_newton(2) - math.Sqrt(2)))
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func sqrt_newton(x float64) float64 {
	// Start with initial guess for z
	z := 1.0

	// Loop a max of 10 times
	for i := 0; i < 10; i++ {
		// Newton's method to update z
		z -= (z*z - x) / (2 * z)
	}

	return z
}
