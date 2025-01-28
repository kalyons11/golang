package main // package main is special, it defines a standalone executable program

// The factored import statement allows multiple packages to be imported in a single import declaration
import (
	"fmt" // fmt is a standard library package that provides formatted I/O
	"math"

	"rsc.io/quote" // quote is a custom package that provides a function that returns a string
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	fmt.Println(math.Pi) // math is a standard library package that provides basic constants and mathematical functions
	// Pi is an exported constant in the math package
}
