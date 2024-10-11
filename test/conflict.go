package main

import (
	"fmt"
)

func main() {
	fmt.Println("Conflict!")
	fmt.Println("Merge Conflict!")

	// root function call
	root()
}

func root() {
	fmt.Println("This is IMPORTANT CODE!")
	fmt.Println("Don't Ignore!")

	// rooted to main
	fmt.Println("This is IMPORTANT CODE!")

}