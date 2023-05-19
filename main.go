package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app = tview.NewApplication()
var menuText = tview.NewTextView()
var box = tview.NewBox()
var flex = tview.NewFlex()
var pages = tview.NewPages()
var flexContainer = tview.NewFlex()

func main() {

	menuText.SetTextColor(tcell.ColorGreen).SetText("(t) to start training\n(q) to quit")
	box.SetBorder(true).SetTitle(" ----- 💪 Welcome to the Math GYM 💪 ----- ")

	flex.SetDirection(tview.FlexRow)

	menu := mainMenu{}
	addMenuList(&menu)

	digits := digitInputs{}
	addDigitsInputForm(&digits)

	ans := answer{}
	calculateAnswer(digits, menu, &ans)

	calc := calculation{}
	addCalculateForm(&calc, menu, ans)

	box.SetDrawFunc(func(screen tcell.Screen, x, y, w, h int) (int, int, int, int) {
		flex.Draw(screen)
		return x, y, w, h
	})

	flexContainer.SetDirection(tview.FlexRow).
		AddItem(flex, 0, 1, true)

	pages.AddPage("Start Menu", menuText, true, true).
		AddPage("Main", flexContainer, true, false)

	pages.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' {
			app.Stop()
		} else if event.Rune() == 't' {
			pages.SwitchToPage("Main")
		} else if event.Rune() == '+' || event.Rune() == '-' || event.Rune() == '*' || event.Rune() == '/' {
			pages.SwitchToPage("Digit Inputs")
		} else if event.Rune() == 'd' {
			pages.SwitchToPage("Calculate")
		}
		return event
	})

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}

type mainMenu struct {
	op string
}

func addMenuList(menu *mainMenu) {
	startText := tview.NewTextView().SetTextColor(tcell.ColorWhite).SetText("Please choose the mathematical operation to perform (add | sub | mul | div):")
	opChoices := tview.NewList().
		AddItem("add", "", '+', func() { menu.op = "+" }).
		AddItem("sub", "", '-', func() { menu.op = "-" }).
		AddItem("mul", "", '*', func() { menu.op = "*" }).
		AddItem("div", "", '/', func() { menu.op = "/" })

	flex.AddItem(startText, 0, 1, false).
		AddItem(opChoices, 10, 1, true)
}

type digitInputs struct {
	firstNumDigits  int
	secondNumDigits int
}

func addDigitsInputForm(digits *digitInputs) {
	digitsInputForm := tview.NewForm()
	digitsInputForm.
		AddInputField("Choose the number of digits of the first operand", "", 20, nil, func(text string) {
			digits.firstNumDigits, _ = strconv.Atoi(text) // handle error
		}).
		AddInputField("Choose the number of digits of the second operand", "", 20, nil, func(text string) {
			digits.secondNumDigits, _ = strconv.Atoi(text) // handle error
		}).
		AddTextView("(d) to continue", "", 0, 0, false, false)

	pages.AddPage("Digit Inputs", digitsInputForm, true, false)
}

type answer struct {
	firstNum  int
	secondNum int
	answer    int
	remainder int
}

func calculateAnswer(digits digitInputs, menu mainMenu, ans *answer) {
	// firstNum, err := generateNumber(digits.firstNumDigits)
	// if err != nil {
	// 	fmt.Println("There was an error generating a number with the specified amount of digits.")
	// 	return
	// }

	// secondNum, err := generateNumber(digits.secondNumDigits)
	// if err != nil {
	// 	fmt.Println("There was an error generating a number with the specified amount of digits.")
	// 	return
	// }

	firstNum := 10
	secondNum := 20

	ans.firstNum = firstNum
	ans.secondNum = secondNum

	switch menu.op {
	case "+":
		ans.answer = firstNum + secondNum
	case "-":
		ans.answer = firstNum - secondNum
	case "*":
		ans.answer = firstNum * secondNum
	case "/":
		ans.answer = firstNum / secondNum
		ans.remainder = firstNum - (ans.answer * secondNum)
	}
}

type calculation struct {
	guess     int
	remainder int
}

func addCalculateForm(calc *calculation, menu mainMenu, ans answer) {
	CalculateForm := tview.NewForm()
	tempText := fmt.Sprintf("Calculate %d %s %d\n", ans.firstNum, menu.op, ans.secondNum)
	CalculateForm.
		AddTextView(tempText, "", 0, 0, false, false).
		AddInputField("answer", "", 20, nil, func(text string) {
			calc.guess, _ = strconv.Atoi(text)
		})

	if menu.op == "/" {
		CalculateForm.
			AddInputField("remainder", "", 20, nil, func(text string) {
				calc.remainder, _ = strconv.Atoi(text)
			})
	}

	pages.AddPage("Calculate", CalculateForm, true, false)
}

func generateNumber(numberOfDigits int) (int, error) {
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
