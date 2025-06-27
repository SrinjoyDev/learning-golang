package main

import "fmt"

func ifElse() {

	if 7%2 == 0 {
		fmt.Println("odd")
	} else {
		fmt.Println("even")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "less than 0")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has MORE THAN  one digit")
	}
}
