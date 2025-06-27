package main

import (
	"fmt"
	"time"
)

func switchCase() {

	i := 2
	fmt.Println("print", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it is a weekday")
	default:
		fmt.Println("it is a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it is not noon yet")
	case t.Hour() > 12:
		fmt.Println("it is past noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("i am a boolean")
		case int:
			fmt.Println("i am a integer")
		default:
			fmt.Println("dont know type ", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hello")

}
