package main

import "fmt"

func main() {
	show_for()
}

func show_for() {
	fmt.Println("Hello from show_for")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// We can create a doubling loop
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
