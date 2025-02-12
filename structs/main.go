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

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X) // 4

	p := &v
	p.X = 1e9      // easier than (*p).X = 1e9
	fmt.Println(v) // {1000000000 2}
}
