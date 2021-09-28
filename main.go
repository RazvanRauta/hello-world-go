package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)

	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')
		response := doctor.Response(userInput)

		if strings.TrimSpace(userInput) == "quit" {
			fmt.Println(response)
			break
		} else {
			fmt.Println(response)
		}

	}

}
