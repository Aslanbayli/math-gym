package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"strconv"
)

func getStrInput() string {
	var input string
	fmt.Scanln(&input)
	return input
}

func getIntInput() (int, error) {
	var input string
	fmt.Scanln(&input)
	return strconv.Atoi(input)
}

func validateDigitsInput() int {
	var numDigits int
	var err error

	for {
		numDigits, err = getIntInput()
		if err == nil && numDigits > 0 {
			break
		}
		fmt.Print("Please enter a valid number of digits: ")
	}
	return numDigits
}

func generateRandomNumber(numberOfDigits int) (int, error) {
	maxLimit := int64(int(math.Pow10(numberOfDigits)) - 1)
	lowLimit := int(math.Pow10(numberOfDigits - 1))

	randomNumber, err := rand.Int(rand.Reader, big.NewInt(maxLimit))
	if err != nil {
		return 0, err
	}
	randomNumberInt := int(randomNumber.Int64())

	// Handling integers between 0, 10^(n-1) .. for n=4, handling cases between (0, 999)
	if randomNumberInt <= lowLimit {
		randomNumberInt += lowLimit
	}

	return randomNumberInt, nil
}

func train() {
	var operation string    // the operation selected by the user (+ | - | * | /)
	var firstNumDigits int  // number of digits of the first operand
	var secondNumDigits int // number of digits of the second operand
	var err error
	var answer int    // correct answer of the operation for checking
	var remainder int // for division operations with remainders
	var guess int     // user's claculation

	validOps := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}

	fmt.Print("Please choose the mathematical operation to perform (+ | - | * | /): ")
	for {
		operation = getStrInput()
		if validOps[operation] {
			break
		}

		fmt.Print("Please enter a valid choice... (+ | - | * | /): ")
	}

	fmt.Print("Choose the number of digits of the first operand: ")
	firstNumDigits = validateDigitsInput()

	fmt.Print("Choose the number of digits of the second operand: ")
	secondNumDigits = validateDigitsInput()

	firstNum, ok := generateRandomNumber(firstNumDigits)
	if ok != nil {
		fmt.Println("There was an error generating a number with the specified amount of digits.")
		return
	}

	secondNum, ok := generateRandomNumber(secondNumDigits)
	if ok != nil {
		fmt.Println("There was an error generating a number with the specified amount of digits.")
		return
	}

	switch operation {
	case "+":
		answer = firstNum + secondNum
	case "-":
		answer = firstNum - secondNum
	case "*":
		answer = firstNum * secondNum
	case "/":
		answer = firstNum / secondNum
		remainder = firstNum - (answer * secondNum)

	}

	fmt.Printf("Calculate %d %s %d\n", firstNum, operation, secondNum)

	for {
		fmt.Print("answer: ")
		guess, err = getIntInput()
		for err != nil {
			fmt.Println("Please make sure your answer is a whole number...")
			fmt.Print("answer: ")
			guess, err = getIntInput()
		}

		if operation == "/" {
			fmt.Print("remainder: ")
			remainderGuess, err := getIntInput()
			for err != nil {
				fmt.Println("Please make sure your answer is a whole number...")
				fmt.Print("remainder: ")
				remainderGuess, err = getIntInput()
			}

			if guess == answer && remainderGuess == remainder {
				fmt.Println("The answer is correct!")
				break
			}
		} else {
			if guess == answer {
				fmt.Println("The answer is correct!")
				break
			}
		}

		fmt.Println("Wrong answer...")
	}

}

func main() {
	fmt.Println("----- 💪 Welcome to the Math GYM 💪 -----")
	for {
		train()
		fmt.Print("\nWould you like to train more? (y | n) ")
		repeat := getStrInput()
		if repeat == "y" {
			continue
		} else if repeat == "n" {
			fmt.Println("\n----- ❤️ Have a good day ❤️ -----")
			fmt.Println("\nPress any key to exit")
			fmt.Scanln()
			break
		} else {
			fmt.Print("Please provide a valid answer (y | n) ")
		}
	}
}
