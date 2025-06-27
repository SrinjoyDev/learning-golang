package main

import "fmt"

//slices are like arraylist from java in golang
//they are typed only by the elements
//uninit slice equals to nul and has length = 0
func slices() {

	var s []string
	fmt.Println("uninit ", s, s == nil, len(s) == 0)
}
