package main

import (
	"fmt"
)

func main() {
	var choice int
	fmt.Print("Enter a number greater than 0: ")
	if _, err := fmt.Scan(&choice); err == nil {
		if choice > 0 {
			for i := 1; i <= choice; i++ {
				fizzbuzz(i)
			}
		} else {
			fmt.Println("That number was not greater than 0. Please try again.")
			main()
		}
	} else {
		fmt.Println("You did not enter a number. Please try again.")
		main()
	}
}

func fizzbuzz(i int) {
	fizz := "fizz"
	buzz := "buzz"

	if i%3 == 0 && i%5 == 0 {
		fmt.Println(i, fizz+buzz)
	} else if i%3 == 0 {
		fmt.Println(i, fizz)
	} else if i%5 == 0 {
		fmt.Println(i, buzz)
	} else {
		fmt.Println(i)
	}
}
