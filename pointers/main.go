package main

import "fmt"

func main() {
	show_pointers()
}

// This isn't too far off from C or C++
func show_pointers() {
	fmt.Println("Hello from show_pointers")

	i, j := 42, 2701
	// point to i
	p := &i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer (dereferencing)
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
