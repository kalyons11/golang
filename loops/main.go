package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	show_for()
}

func show_for() {
	fmt.Println("Hello from show_for")

	// We can use defer to ensure a function call is done at the end of the function
	// This is useful for closing files, connections, etc.
	defer fmt.Println("World (deferred).")

	// Defers are pushed onto a stack and executed in a last-in-first-out order
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	} // This will print 9 to 0

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

	show_switch("linux")
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

func show_switch(os string) {
	fmt.Println("Hello from show_switch")

	// Switch statements in golang are similar to other languages
	// We can use a switch without a condition to create a long if-else chain
	// We can also use a switch with a condition
	// We can also use a switch with a condition and a fallthrough statement
	// We can also use a switch with a condition and a default statement
	// We can also use a switch with a condition and a default statement with a fallthrough statement
	switch os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
		fallthrough // This will print the next case as well
	default:
		fmt.Printf("%s. fallthrough.\n", os)
	}

	// We can use a switch with no condition to create a long if-else chain
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
