// simply shows 1 function calling another with 2 arguments to print out amount of paint needed

package main

import "fmt"

func paintNeeded(width float64, height float64) {
	area := width * height
	fmt.Printf("%.2f liters needed\n", area/10.0) // one way to print formatted results
}

func main() {
	paintNeeded(4.2, 3.0)
}
