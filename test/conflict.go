package main

import (
	"fmt"
)

func main() {
<<<<<<< HEAD
	fmt.Println("Merge Conflict!")
=======
	fmt.Println("Conflict!")
>>>>>>> conflicts

	// root function call
	root()
}

func root() {
<<<<<<< HEAD
	// rooted to main
	fmt.Println("This is IMPORTANT CODE!")
=======
	fmt.Println("This is IMPORTANT CODE!")
	fmt.Println("Ignore!")
	fmt.Println("Don't Ignore!")
    fmt.Println(repo(45))
}

// exp code 
func repo(v int) int{
    return v+=2;
>>>>>>> conflicts
}
