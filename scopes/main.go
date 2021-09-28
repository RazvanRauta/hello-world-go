package main

import "fmt"

// package level var
var one = "One"
var PublicVar = "This is a public var"

func main() {

	var one = "this is a black level var"

	fmt.Println(one)
	myFunc()
}

func myFunc() {
	fmt.Println(one)
}

func ExportedFunc() {
	fmt.Println("This function is public")
}
