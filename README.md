# Math GYM

This console app is a personal improvement project designed to enhance mathematical skills by providing a customizable environment where users can select a mathematical operator and generate algebraic expressions with a chosen number of digits, allowing practice and improvement in performing mental calculations without relying on a calculator.

It is written entirely in Go programming language. Initially I was planning on adding some terminal UI but eventually decided to keep it as simple as possible.


# Compilation and Runing

First and foremost, make sure to have `Go` insattled on your system. If you simply want to run the code in your terminal you can simply type this in your terminal:

```bash
make run
```

This will build and run the code.

However, if you wish to build and create an executable instead, run one of the two commands:

### For windows
```bash
make build-win
```

### For MacOS or Linux
```bash
make build-unix
```

## App Usage

Once you run the code it will open a prompt in the terminal asking you to choose a mathematical expressions

img

Subsequently the program will open two more prompts to input the digits of two numbers for generting an algebraic expression. Keep in mind that these are the digit sizes not numbers themselves, the numbers are generated randomly each time. 

img

If you answer correctly you can keep using the app.

img

Otherwise, you will have to give a correct answer until you can proceed further.

img

## Future plans

In the future, I plan on adding a timer functionality so that users can measure their imporvement based on how much faster they were able to compute an expressions.

