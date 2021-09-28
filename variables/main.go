package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and don't type your number in, just press ENTER when ready."

var reader = bufio.NewReader(os.Stdin)

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// rand generates a number between 0 and whatever is passed as a parameter
	// we add 2 to it because we want the number used to be at least 2, and no
	// greater than 10 (multiplying by 1 makes the game a bit silly)
	firstNumber := rand.Intn(8) + 2
	secondNumber := rand.Intn(8) + 2
	subtraction := rand.Intn(8) + 2
	answer := firstNumber*secondNumber - subtraction

	playTheGame(firstNumber, secondNumber, subtraction, answer)
}

func playTheGame(firstNumber, secondNumber, subtraction, answer int) {

	fmt.Printf(`
	Guess the Number Game
	---------------------

	Think of a number between 1 and 10
	%s
	`, prompt)
	reader.ReadString('\n')

	// take them through the games
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')

	// give them the answer
	fmt.Println("The answer is", answer)
}
