package main

import (
	"fmt"
	"math/rand"
)

func RandomNumber(n int) {
	x := rand.Intn(n)
	fmt.Printf("the value of %x is \t", x)
	if x >= 100 && x <= 200 {
		fmt.Println("the number between 100 and 200")
	} else if x >= 1 && x <= 100 {
		fmt.Println("the number between 1 to 100")
	}

}
