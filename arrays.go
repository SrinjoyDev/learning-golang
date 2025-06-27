package main

import "fmt"

func arrays() {

	//defining an array with size 5 , by default values are 0
	var a [5]int
	a[4] = 100
	fmt.Println("set ", a)
	fmt.Println("get ", a[4])

	fmt.Println("len ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl : ", b)

	//make the compiler count the number of elements in the array
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl :", b)

	//: elements in between will be zeroo
	var c = [...]int{1, 2: 6, 7, 8}
	fmt.Println(" : use ", c)

	//2d arrays
	var twoDarray [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoDarray[i][j] = i + j
		}
	}
	fmt.Println("2d : ", twoDarray)

	//we can decalre like this for 2d arrays>
	var twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("self defining 2d arrays : ", twoD)

}
