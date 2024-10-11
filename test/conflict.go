package main

import (
	"fmt"
)

func main() {
	fmt.Println("Merge Conflict!")
	fmt.Println("Conflict!")
	fmt.Println("This is IMPORTANT CODE!")
	fmt.Println("Don't Ignore!")
    fmt.Println(repo(45))
	// root function call
	root()
}

// exp code 
func repo(v int) int{
    return v+=2;
}
