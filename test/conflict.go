package main

import (
	"fmt"
)

func main() {
	fmt.Println("Conflict!")

	// root function call
	root()
}

func root() {
	fmt.Println("This is IMPORTANT CODE!")
	fmt.Println("Ignore!")
	fmt.Println("Don't Ignore!")
}
