package main

import "fmt"

func main() {
	demo_structs()
}

// A struct is a collection of fields
type Vertex struct {
	X, Y int
}

func demo_structs() {
	fmt.Println("Hello from demo_structs")
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X) // 4

	p := &v
	p.X = 1e9      // easier than (*p).X = 1e9
	fmt.Println(v) // {1000000000 2}

	// Example of struct literals
	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		p2 = &Vertex{1, 2} // has typ2e *Vertex
	)

	fmt.Println(v1, p2, v2, v3) // {1 2} &{1 2} {1 0} {0 0}
}
