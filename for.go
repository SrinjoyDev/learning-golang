package main

import "fmt"

func forLoop() {

	//normal for loop >>
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//rnage -> loop over 3 times
	for i := range 3 {
		fmt.Println(i)
	}

	var j int = 1
	for {
		fmt.Println(j)
		if j == 5 {
			break
		}
		j = j + 1
		continue
	}

	//6 times loop for even number continue otherwise priunt
	//that is print the odd nos
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
