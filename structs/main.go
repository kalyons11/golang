package main

import "fmt"

func main() {
	demo_structs()
}

// A struct is a collection of fields
type Vertex struct {
	X int
	Y int
}

func demo_structs() {
	fmt.Println("Hello from demo_structs")
	fmt.Println(Vertex{1, 2})
}
