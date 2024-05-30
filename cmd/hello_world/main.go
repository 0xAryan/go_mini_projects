package main

import (
	"fmt"
	"os"
)

func main() {
    var name string
    var age int
    fmt.Print("Enter your name and age: ")
    _, err := fmt.Scanf("%v  |   %v", &name, &age)
	if err != nil {
		fmt.Fprint(os.Stderr, err, "\n")
	}
    fmt.Printf("Name: %s, Age: %d\n", name, age)
}
