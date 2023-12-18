package main

import (
	"fmt"
)

func main() {
	fmt.Print("hello rajendra")
	//s1 := puppy.Bark()
	//s2 := puppy.BigBark()
	//fmt.Println(s1)
	//fmt.Println(s2)
	//nestedLoop()
	//randomNumber(10)
	//printSlice()
	rangeSlice()
}

func nestedLoop() {
	x := 5
	for i := 0; i < x; i++ {
		for j := 0; j < x; j++ {
			fmt.Printf("here is the nested loop %d and %d", j, i)
		}
	}
}

func printIntSlcie() {
	xi := []int{1, 2, 3, 4, 5, 6}
	for i, v := range xi {
		fmt.Printf("value of the %d and %d \n", i, v)

	}
}

func playWithArray() {
	a := [2]int{1, 3}
	fmt.Println(a)

	b := [...]string{"raj", "ins", "tiger"}
	fmt.Println(b)

}

func printSlice() {
	xs := []string{"rajendra", "spro", "tt", "tt"}

	for i, v := range xs {
		fmt.Printf("value : %d  index : %s\n", i, v)
	}
}

func rangeSlice() {
	// Creating a slice
	names := []string{"rajendra", "spro", "tt", "tt"}

	//I don't 1 value or i dom't index
	for _, value := range names {
		fmt.Printf(" Value: %s\n", value)
	}

	fmt.Println("this is println the lenght of my index", len(names))
}
