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
    fmt.Println(repo(45))
}

// exp code 
func repo(v int) int{
    return v+=2;
}
