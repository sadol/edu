package main

import (
	//"./genstack"
	"./polish"
	"bufio"
	"fmt"
	"os"
)

func main() {
	const QUIT = "q"
	inputReader := bufio.NewReader(os.Stdin)
	var usersExpression string
	var err error
	var myNPN *polish.NPN
	var result polish.Operand

	//testo()
	for {
		fmt.Println("Enter Normal Polish Notation expression or `q' to quit:")
		if usersExpression, err = inputReader.ReadString('\n'); err != nil {
			fmt.Println("Input error.")
			return
		}
		// remove \n from the input
		usersExpression = usersExpression[:len(usersExpression)-1]
		if usersExpression == "q" {
			fmt.Println("Bye.")
			return
		}
		if myNPN, err = polish.NewNPN(usersExpression); err != nil {
			fmt.Println("NewNPN error.")
			return
		}
		if result, err = myNPN.Calculate(); err != nil {
			fmt.Println("Calculate error.")
			return
		} else {
			fmt.Printf("Expression:<%s>, gave:<%d>.\n", usersExpression, result)
		}
	}
}

func testo() {
	var expression1 = "+ 3 3" // 3+3 = 6
	var result12 = 6
	var expression2 = "+ / - * 3 4 2 5 4"     // 3+(((3*4)-2)/5) = 6
	var expression3 = "- - 8 1 * / 4 2 + 3 3" // (8-1) - ((4/2) * (3+3)) = -5
	var result polish.Operand
	var result3 = -5
	var err error
	var myNPN *polish.NPN

	if myNPN, err = polish.NewNPN(expression1); err != nil {
		return
	}
	if result, err = myNPN.Calculate(); err != nil {
		return
	} else {
		fmt.Printf("Expression:<%s>, should give:<%d>, gave:<%d>.\n",
			expression1, result12, result)
	}
	if myNPN, err = polish.NewNPN(expression2); err != nil {
		return
	}
	if result, err = myNPN.Calculate(); err != nil {
		return
	} else {
		fmt.Printf("Expression:<%s>, should give:<%d>, gave:<%d>.\n",
			expression2, result12, result)
	}
	if myNPN, err = polish.NewNPN(expression3); err != nil {
		return
	}
	if result, err = myNPN.Calculate(); err != nil {
		return
	} else {
		fmt.Printf("Expression:<%s>, should give:<%d>, gave:<%d>.\n",
			expression3, result3, result)
	}
}
