package main

import "fmt"

func main() {
	amount := 6
	double2(amount)
	fmt.Println(amount) // this shows only 6 because it is a pass by value and the double2 function only alters the copy of the parameter
}

func double2(number int) {
	number *= 2
	//fmt.Println(number) when we print here we get 12
}
