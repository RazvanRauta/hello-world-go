package main

import "fmt"

func variables() {
	// * one -way - declare, then assign (two steps)

	var number int
	number = 2

	// * another way. declare type and name
	var secondNumber = 5

	// one step variable
	thirdNumber := 7

	// * print formated
	fmt.Printf("number: %d, secondNumber: %d, thirdNumber %d \n",
		number,
		secondNumber,
		thirdNumber)

}
