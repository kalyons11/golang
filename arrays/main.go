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

	// We can define a slice literal, which is like an array literal without the length
	var slice = []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	// We can use the make function to create a slice
	var make_slice = make([]int, 5)
	fmt.Println(make_slice)

	// We can create a slice of structs
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}
	fmt.Println(people)

	list := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(list)

	// Slice defaults
	// When slicing, you may omit the high or low bounds to use their defaults instead.
	// The default is zero for the low bound and the length of the slice for the high bound.
	// For the array
	// var a [10]int
	// these slice expressions are equivalent:
	// a[0:10]
	// a[:10]
	// a[0:]
	// a[:]

	test := []int{2, 3, 5, 7, 11, 13}

	test = test[1:4]
	fmt.Println(test)

	test = test[:2]
	fmt.Println(test)

	test = test[1:]
	fmt.Println(test)
}
