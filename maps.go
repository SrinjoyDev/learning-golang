package main

import (
	"fmt"
	"maps"
)

func hashMaps() {

	//map[key-type]val type
	m := make(map[string]int)

	//set key value pairs>>
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("m ", m)

	//get a value with the key name
	val := m["k1"]
	val2 := m["k2"]
	fmt.Println("values gotten for the keys k1 and k2 : ", val, val2)

	fmt.Println("map value :", len(m))

	//delete key>
	delete(m, "k2")
	fmt.Println("deleted from the map k2 ", m)

	//clear all the key val pairs from the map >>
	clear(m)
	fmt.Println("cleared map ", m)

	//blank identifier when u want to ignore some value we use a blank identifier.
	_, prs := m["k2"]
	fmt.Println("prs:", prs) //false

	//declare a map and init the values in the same line like this >
	n := map[string]int{
		"foo": 1,
		"bar": 2,
	}

	n2 := map[string]int{
		"foo": 1,
		"bar": 2,
	}

	//this blank identifier is used when we only care if the value exisits and dont care about what the value of the key is
	_, prs2 := n["bar"]
	fmt.Println("experimenting blank identifier : ", prs2) //true

	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}

}
