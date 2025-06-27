package main

import "fmt"

func variables() {
	var a = "intital"
	fmt.Println(a)
	var b, c int = 1, 2
	fmt.Println(b, c)
	var d int
	fmt.Println(d) //default val is 0

	//this method is only applicable inside functions
	f := "apple" //same as var f string = "apple"
	fmt.Println(f)
}
