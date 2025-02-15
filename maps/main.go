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
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}
	fmt.Println(n)

	// We can insert, delete, retrieve and check if a key is present in a map
	n["Apple"] = Vertex{37.7749, -122.4194}
	fmt.Println(n)

	delete(n, "Apple")
	fmt.Println(n)

	_, ok := n["Apple"]
	fmt.Println("Apple is present in the map:", ok)

	_, ok = n["Google"]
	fmt.Println("Google is present in the map:", ok)

	// Check a non-existent key
	_, ok = n["Microsoft"]
	fmt.Println("Microsoft is present in the map:", ok) // false

	// We can iterate over a map
	for key, value := range n {
		fmt.Printf("Key: %s, Value: %v\n", key, value)
	}

	// We can also iterate over a map by only getting the keys
	for key := range n {
		fmt.Println("Key:", key)
	}

	// We can also iterate over a map by only getting the values
	for _, value := range n {
		fmt.Println("Value:", value)
	}

	// We can also create a map of maps
	var o = map[string]map[string]Vertex{
		"USA": {
			"California": {37.7749, -122.4194},
			"New York":   {40.7128, -74.0060},
		},
		"India": {
			"Delhi":  {28.7041, 77.1025},
			"Mumbai": {19.0760, 72.8777},
		},
	}
	fmt.Println(o)

	// We can also delete a key from a map of maps
	delete(o["USA"], "California")
	fmt.Println(o)
}
