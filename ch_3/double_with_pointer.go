package main

import "fmt"

func main() {
	amount := 6
	double3(&amount) //Pass a pointer instead of the variable value
	fmt.Println(amount)
}

func double3(number *int) { //Accept a pointer instead of an int value.
	*number *= 2 //update the value at the pointer
}
