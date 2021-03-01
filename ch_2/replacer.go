package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G# r#cks!"
	/*
		Ths is a block of comments.
		That spans multiple lines.
	*/
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
}
