package main

import (
	"fmt"
	"math/rand"
)

func RandomNumber(n int) int {
	x := 40
	fmt.Printf("the value of %x is \t", x)

	if z := x * rand.Intn(n); z >= x {
		fmt.Println("the number between 100 and 200")
	} else if x >= 1 && x <= 100 {
		fmt.Println("the number between 1 to 100")
	}
	return x
}

func offSet(t int) int {
	if x := true; x {
		return t
	}
	return t
}

func testSwitchCase(x int) {
	switch x {
	case 1:
		fmt.Println("rajendra")

	}
}
