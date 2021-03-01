package main

import (
	"fmt"
)

func main() {
	var width, height, area float64
	width = 4.2
	height = 3.5
	area = width * height
	fmt.Println(area/10.0, "liters needed")
	width = 5.2
	height = 4.5
	area = width * height
	fmt.Printf("%.2f liters needed\n", area/10)
}
