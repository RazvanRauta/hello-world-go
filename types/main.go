package main

import (
	"fmt"
	"log"
	"sort"
)

// basic types (numbers, strings, booleans)
var myInt int
var myUint uint
var myFloat float32
var myFloat64 float64

// aggregate types (array, struct)

type Car struct {
	NumberOfTires int
	Luxury        bool
	BucketSeats   bool
	Make          string
	Model         string
	Year          int
}

// reference types (pointers, slices, maps, functions, channels)

// interface type

func main() {
	// basic types (numbers, strings, booleans)
	myInt = 10
	myUint = 20

	myFloat = 10.1
	myFloat64 = 100.1

	log.Println(myInt, myUint, myFloat, myFloat64)

	myString := "Razvan"

	log.Println(myString)

	myString = "Raz"

	var myBool = true
	myBool = false

	log.Println(myBool)

	// End of basic types --------

	// aggregate types (array, struct)

	var myStrings [3]string

	myStrings[0] = "cat"
	myStrings[1] = "dog"
	myStrings[2] = "mouse"

	fmt.Println("First element is array is", myStrings[0])

	myCar := Car{
		NumberOfTires: 4,
		Luxury:        true,
		BucketSeats:   true,
		Make:          "Dacia",
		Model:         "Logan",
		Year:          2005,
	}

	fmt.Printf("My car is a %d %s %s \n", myCar.Year, myCar.Make, myCar.Model)

	// End of aggregate types --------

	// reference types (pointers, slices, maps, functions, channels)

	// ---
	x := 10

	myFirstPointer := &x // pointer to address in memory

	fmt.Println("x is", x)
	fmt.Println("myFirstPointer is", myFirstPointer)

	*myFirstPointer = 15 // add new value to the address in memory

	fmt.Println("x is now", x)

	changeValueOfPointer(&x)

	fmt.Println("After function call, x is now", x)

	// ---

	var animals []string // slice

	animals = append(animals, "dog")
	animals = append(animals, "fish")
	animals = append(animals, "cat")
	animals = append(animals, "horse")

	fmt.Println(animals)

	for i, x := range animals {
		fmt.Println(i, x)
	}

	fmt.Println("Element 0 is", animals[0])

	fmt.Println("First two elements are", animals[0:2])        // print first 2 elements
	fmt.Println("The slice is", len(animals), "elements long") // length

	//sort slice
	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))

	sort.Strings(animals)
	fmt.Println(animals)

	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))

	animals = deleteFromSlice[[]string](animals, 1, 2)
	fmt.Println(animals)

}

func changeValueOfPointer(num *int) {
	*num = 25
}

func deleteFromSlice[T any](s []string, i, j int) []string {
	return append(s[:i], s[j:]...)
}
