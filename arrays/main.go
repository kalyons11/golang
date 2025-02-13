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

	// We can also declare and initialize an array in one line
	var primes = [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// We can create a slice from an array
	var s []int = primes[1:4]
	fmt.Println(s) // [3 5 7]

	// Slices are just references to arrays
	s[0] = 100
	fmt.Println(primes) // [2 100 5 7 11 13]
}
