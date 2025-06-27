package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func constants() {

	fmt.Println(s)

	const n = 5000000000 //numeric constant has no type until it is given one
	const d = 3e20 / n   //constant with no type still

	fmt.Println(int64(d)) //converted go into a int 64
	//we didnt assign n any type so when we pass n to math.sin it automativally converts n to float
	fmt.Println(math.Sin(n)) // n should be a float the return val is also a float , not permanently but does that for the operation

}
