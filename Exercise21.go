package main

import (
	"fmt"
	"math/rand"
)

func init(){
	fmt.Println(" I am in the main ")
}

func random(){
	x:= rand.Intn(2);
	y:=rand.Intn(1);
	if x<4 && y<4
	fmt.Println("here is the value of x and y")
	else if x>4 && y> 4 
	fmt.Println("here is the values of x and y")

}

func checkForLoop(){
	for i=0;i<=100;i++ {
		fmt.Printf("here is the number %d", i);
	}

	for x<10{
		fmt.Println("infinit loop")
		x++;
	}

	for {
		if x>20 {
			fmt.Println("this the thrid way to declare the for loop")
			break;
		}
		x++;
	}

}

func infinitLoop(){

	x:=20

	for {
		if x>=20 {
			fmt.Printf("here is the line %d",x)
			break;
		}	
		x--

	}
	
}

func printOdd(){
	 x:=20
	for i:=0;i<x;i++ {
		if i%2!=0 {
			println("the older number")
		}
	}
}

