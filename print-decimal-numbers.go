package main

// Import packages
import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	// Finding a Random number using rand.Intn package
	fmt.Println("Random Number is: ", rand.Intn(10))
	// Don't need to add '\n' at the end to break the line.

	// Embed values in string - (embedding square root of 7)
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	// We want to add '\n' at the end if we needed to break the line
	// Embed variables or values in string using different place holders like below

	// %g - floating-point numbers

	// %f - It is also using to print floating-point. but we can limit the floating point like %.2f for 2 decimal point
	fmt.Printf("Now you %.2f\n", 3.222)

	// %e / % E  - For scientific notation. use for extremely large or small numbers.
	// Example: fmt.Printf("%e", 123400000.0) → 1.234000e+08

	// %g / %G  - This is the most flexible and common choice. It automatically
	// switches between %f and %e based on which one produces the shortest,
	// most compact output. Crucially, it removes unnecessary trailing zeros.

}
