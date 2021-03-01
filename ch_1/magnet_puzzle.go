package main

import (
	"fmt"
)

func main() {
	    var originalCount int = 10
	    var eatenCount int = 6

	    fmt.Println("I started with", originalCount, "apples.")
	    fmt.Println("Some jerk at", eatenCount, "apples.")
	    fmt.Println("There are", originalCount-eatenCount, "apples left.")
}