package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
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
	var opStr string        // the operation selected by the user (add, sub, mul, div)
	var firstNumDigits int  // number of digits of the first operand
	var secondNumDigits int // number of digits of the second operand
	var err error
	var op string     // operation in symbol form (+, -, *, /)
	var answer int    // correct answer of the operation for checking
	var remainder int // for division operations with remainders
	var guess int     // user's claculation

	validOps := map[string]bool{
		"add": true,
		"sub": true,
		"mul": true,
		"div": true,
	}

	fmt.Print("Please choose the mathematical operation to perform (add | sub | mul | div): ")
	for {
		opStr = getStrInput()
		opStr = strings.ToLower(opStr)
		if validOps[opStr] {
			break
		}

		fmt.Print("Please enter a valid choice... (add | sub | mul | div): ")
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

	switch opStr {
	case "add":
		op = "+"
		answer = firstNum + secondNum
	case "sub":
		op = "-"
		answer = firstNum - secondNum
	case "mul":
		op = "*"
		answer = firstNum * secondNum
	case "div":
		op = "/"
		answer = firstNum / secondNum
		remainder = firstNum - (answer * secondNum)

	}

	fmt.Printf("Calculate %d %s %d\n", firstNum, op, secondNum)

	for {
		fmt.Print("answer: ")
		guess, err = getIntInput()
		for err != nil {
			fmt.Println("Please make sure your answer is a whole number...")
			fmt.Print("answer: ")
			guess, err = getIntInput()
		}

		if op == "/" {
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
			fmt.Scanln()
			break
		} else {
			fmt.Print("Please provide a valid answer (y | n) ")
		}
	}
}
