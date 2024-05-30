package main

import (
	"fmt"
)


func main() {
	var i int
	fmt.Print("Enter a no between 1 and 10 : ")
	fmt.Scan(&i)
	if i >= 1 && i <= 10 {
		fmt.Println("It's between 1 to 10")
	} else {
		fmt.Println("It out of range")
	}
}