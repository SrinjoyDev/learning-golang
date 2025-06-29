package main

import "fmt"

func vals() (int, int) {
	return 3, 5
}

func multiplereturnvalues() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// if we only want the subset of the return values then we do like this
	_, c := vals()
	fmt.Println("subset of the return values wuth a blank identifier", c)
}
