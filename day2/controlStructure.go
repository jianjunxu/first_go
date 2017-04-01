package main

import (
	"fmt"
)

func main() {
	// if
	if x := 20; x > 10 {
		fmt.Println("x is greater than 10.")
	} else {
		fmt.Println("x is less than 10.")
	}
	// goto
	myGoto()
	// for
	for i := 0; i < 5; i++ {
		fmt.Printf("for:%d;", i)
	}
	fmt.Println()

	arr := [...]int{1, 2, 3}
	for i, value := range arr {
		fmt.Printf("arr[%d]:%d;", i, value)
	}
	fmt.Println()
	// switch
	iSwitch := 10
	switch iSwitch {
	case 1:
		fmt.Println("i is equals to 1.")
	case 2, 3, 4:
		fmt.Println("i is equals to 2,3 or 4.")
	case 10:
		fmt.Println("i is equals to 10.")
	default:
		fmt.Println("i is not match any value.")
	}
}

func myGoto() {
	i := 0
	Here:
	fmt.Printf("goto:%d;", i)
	i++
	if (i > 3) {
		fmt.Println()
		return
	}
	goto Here
}


