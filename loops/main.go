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
}
