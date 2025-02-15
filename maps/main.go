package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex // map[key_type]value_type

func main() {
	run_maps()
}

func run_maps() {
	fmt.Println("Hello from run_maps")

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// We can have map literals
	var n = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(n)
}
