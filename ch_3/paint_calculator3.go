// shows a function having 2 returns and error handling

package main

import (
	"fmt"
	"log"
)

func paintNeeded3(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	area := width * height
	return area / 10.0, nil
}

func main() {
	amount, err := paintNeeded3(4.2, -3.0)
	if err != nil {
		// fmt.Println(err) // this would simply print the error
		log.Fatal(err) // this displays the error and also exits the program
	} else {
		fmt.Printf("%0.2f liters needed\n", amount)
	}
}
