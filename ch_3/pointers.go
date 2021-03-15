package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt int = 4
	var myIntPointer *int // declare a variable that holds a pointer to an int
	myIntPointer = &myInt // Assign a pointer to the variable.

	fmt.Println(reflect.TypeOf(&myInt))
	fmt.Println(*myIntPointer) //printing the value of the pointer
	fmt.Println(myIntPointer)  //printing the pointer itself

	var myFloat float64 = 98.6
	var myFloatPointer *float64
	myFloatPointer = &myFloat

	fmt.Println(*myFloatPointer)
	fmt.Println(myFloatPointer)
	fmt.Println(reflect.TypeOf(&myFloat))
	var myBool bool = true
	myBoolPointer := &myBool //A short declaration for a pointer variable.

	fmt.Println(*myBoolPointer)
	fmt.Println(myBoolPointer)
	fmt.Println(reflect.TypeOf(&myBool))
}
