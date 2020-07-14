package main

import (
	"fmt"
	"math/rand"
	"time"
)

type NumberType func(int, int) int

type RandomNumber struct {
	minNum int
	maxNum int
	attempts int
	number NumberType
}

func main() {
	random := RandomNumber{
		minNum: 1,
		maxNum: 20,
		attempts: 5,
		number: func(maxNum int, minNum int) int {
			rand.Seed(time.Now().UnixNano())
			return rand.Intn(maxNum - minNum) + minNum
		},
	}
	secretNumber := random.number(random.maxNum, random.minNum)
	tries := 0

	fmt.Println("Guess number between ", random.minNum ," and ", random.maxNum)

	for tries < random.attempts  {
		fmt.Println(tries)
		var guess intx
		fmt.Scanln(&guess)

		fmt.Println("Your guess was ", guess)

		if (guess == secretNumber) {
			fmt.Println("You won :)")
			break
		} else if (guess < secretNumber) {
			fmt.Println("Guess bigger number!")
			tries += 1
		} else if (guess > secretNumber) {
			fmt.Println("Guess smaller number!")
			tries += 1
		}
	}

	fmt.Println("\nYou ran out of tries. \nBetter luck next time.")
}
