package main

import "fmt"

func pointer() {
	var myString string
	myString = "rajendra"
	fmt.Print(myString)
	changeUsingPointer(&myString)
}

func changeUsingPointer(s *string) {
	newValue := "rajendra"
	*s = newValue
}
