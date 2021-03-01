// shows a function returning a value to be used instead of just printed.
// To declare that a function returns a value, add the type of that return value following the parameters in the function declaration

package main

import (
	"fmt"
)

func paintNeeded2(width float64, height float64) float64 {
	area := width * height
	return area / 10.0
}

func double(number float64) float64 { // example of a function that returns something, notice the type specified
	return number * 2
}

func main() {
	dozen := double(6.0)
	fmt.Println(dozen)
	fmt.Println(double(4.2))

	var amount, total float64
	amount = paintNeeded2(4.2, -3.0)
	fmt.Printf("%0.2f liters needed\n", amount)
	total += amount
	amount = paintNeeded2(5.2, 3.5)
	fmt.Printf("%0.2f liters needed\n", amount)
	total += amount
	fmt.Printf("Total: %0.2f liters\n", total)
}
