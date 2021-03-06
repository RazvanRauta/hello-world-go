package main

import (
	"fmt"
	"log"
	"sort"
	"unicode"

	"github.com/eiannone/keyboard"
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
	MaxSpeed      int
}

// ! liking a function to type
func (a *Car) Speed() {
	fmt.Printf("A %s goes with %d \n", a.Make, a.MaxSpeed)
}

func (a *Car) Whatyear() {
	fmt.Printf("The %s is from %d \n", a.Make, a.Year)
}

// ! channels

var keyPressChan chan rune

func main() {
	// ! basic types (numbers, strings, booleans)
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

	// ! End of basic types --------

	// ! aggregate types (array, struct)

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

	// ! End of aggregate types --------

	// * reference types (pointers, slices, maps, functions, channels)

	// ! ---
	x := 10

	myFirstPointer := &x // ? pointer to address in memory

	fmt.Println("x is", x)
	fmt.Println("myFirstPointer is", myFirstPointer)

	*myFirstPointer = 15 // ? add new value to the address in memory

	fmt.Println("x is now", x)

	changeValueOfPointer(&x)

	fmt.Println("After function call, x is now", x)

	// ! ---

	var animals []string // ? slice

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

	// ? sort slice
	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))

	sort.Strings(animals)
	fmt.Println(animals)

	fmt.Println("Is it sorted?", sort.StringsAreSorted(animals))

	animals = deleteFromSlice(animals, 1, 2)
	fmt.Println(animals)

	// * maps ----------

	intMap := make(map[string]int) // ! Record<string,int> not sorted

	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	for key, value := range intMap {
		fmt.Println(key, value)
	}

	delete(intMap, "four")

	// ? check if element is in map
	el, ok := intMap["four"]

	if ok {
		fmt.Println(el, "is in map")
	} else {
		fmt.Println("Element is not in map")
	}

	intMap["two"] = 10

	for key, value := range intMap {
		fmt.Println(key, value)
	}

	myTotal := sumMany(2, 3, 4, 5)

	fmt.Println("my total is", myTotal)

	// ! functions

	myNewCar := Car{
		NumberOfTires: 4,
		Luxury:        true,
		BucketSeats:   true,
		Make:          "BMW",
		Model:         "435",
		Year:          2022,
		MaxSpeed:      250,
	}

	myNewCar.Speed()
	myNewCar.Whatyear()

	fmt.Println(" --------- ")

	myOldCar := Car{
		NumberOfTires: 4,
		Luxury:        false,
		BucketSeats:   false,
		Make:          "Dacia",
		Model:         "Logan",
		Year:          2005,
		MaxSpeed:      160,
	}
	myOldCar.Speed()
	myOldCar.Whatyear()

	// ! channels

	keyPressChan = make(chan rune)

	go listenForKeyPress()

	fmt.Println("Press any key, or q to quit")
	_ = keyboard.Open()
	defer func() {
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey()
		if unicode.ToLower(char) == 'q' {
			break
		}

		keyPressChan <- char
	}

}

// ! functions

func changeValueOfPointer(num *int) {
	*num = 25
}

func deleteFromSlice(s []string, i, j int) []string {
	return append(s[:i], s[j:]...)
}

// * Variadic functions can be called with any number of trailing arguments
func sumMany(nums ...int) int {
	total := 0

	for _, x := range nums {
		total = total + x
	}

	return total
}

// ! Channels

// listenForKeyPress is called using the "go" keyword, so it runs as a goroutine while the
// calling function (main) continues to execute
func listenForKeyPress() {
	for {
		key := <-keyPressChan
		fmt.Println("You pressed", string(key))
	}
}
