/**
 *  @author: Razvan Rauta
 *  Date: Oct 05 2021
 *  Time: 12:08
 */

package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	coffees := map[int]string{
		1: "Cappuccino",
		2: "Latte",
		3: "Americano",
		4: "Mocha",
		5: "Macchiato",
		6: "Espresso",
	}

	printMenu()

	for {
		char, _, err := keyboard.GetSingleKey()

		if err != nil {
			log.Fatal(err)
		}

		if strings.ToLower(string(char)) == "q" {
			fmt.Println("Bye!!")
			break
		}

		i, _ := strconv.Atoi(string(char))

		if _, ok := coffees[i]; ok {
			fmt.Printf("You chose %s\n", coffees[i])
		} else {
			fmt.Printf("Option %q not found, please select one from the menu:", char)
			printMenu()
		}
	}

	fmt.Println("Program exiting..")

}

func printMenu() {
	fmt.Println(`
	MENU
	------
	1 - Cappuccino
	2 - Latte
	3 - Americano
	4 - Mocha
	5 - Macchiato
	6 - Espresso
	Q - Quit the program
	`)
}
