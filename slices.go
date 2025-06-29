package main

import (
	"fmt"
	"slices"
)

// slices are like arraylist from java in golange
// they are typed only by the elements
// uninit slice equals to nul and has length = 0
func slicesLearn() {

	var s []string
	fmt.Println("uninit ", s, s == nil, len(s) == 0)

	//make the slice initially of length 3 , it can grow with size
	s = make([]string, 3)
	fmt.Println("emp slice : ", s, len(s), cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set", s)
	fmt.Println("get ", s[2])

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("after appending d e f ", s)
	fmt.Println("len ", len(s))

	//copying a slice by creating an empty slic of the same length as s and coptu into c from s
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("new slice after coptied from slice s : -", c)

	//get in range from a slice
	fmt.Println(s[2:5])
	fmt.Println(s[2:]) //2 onwards print

	//declare and init a slice in a single line
	t := []string{"g", "h", "i"}
	fmt.Println("t ", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

}
