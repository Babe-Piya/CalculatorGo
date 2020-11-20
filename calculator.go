package main

import (
	"fmt"
	"os"
	"strconv"
)

var first string
var second string
var operator string
var result int64

func main() {

	if len(os.Args) > 3 {
		operator = operatorCheck(os.Args[1])
		numFirst, err := operandCheck(os.Args[2])
		if err != nil {
			fmt.Println("Invalid number Please try again.")
		}
		numSecond, err := operandCheck(os.Args[3])
		if err != nil {
			fmt.Println("Invalid number Please try again.")
		}
		cal(operator, numFirst, numSecond)
	}
	for {
		fmt.Print("Please enter operator :  ")
		fmt.Scan(&operator)
		operator = operatorCheck(operator)

		fmt.Print("Please enter first number : ")
		fmt.Scan(&first)
		numFirst, err := operandCheck(first)
		if err != nil {
			fmt.Println("Invalid number Please try again.")
			continue
		}

		fmt.Print("Please enter second number :  ")
		fmt.Scan(&second)
		numSecond, err := operandCheck(second)
		if err != nil {
			fmt.Println("Invalid number Please try again.")
			continue
		}

		cal(operator, numFirst, numSecond)
	}

}

func cal(operator string, numFirst int64, numSecond int64) {
	var result int64

	switch operator {
	case "+":
		result = numFirst + numSecond
	case "-":
		result = numFirst - numSecond
	case "*":
		result = numFirst * numSecond
	case "/":
		result = numFirst / numSecond
	}
	fmt.Printf("Result : %d\n", result)
}

func operatorCheck(operator string) string {

	for i := 0; i < 1; {
		switch operator {
		case "+", "-", "*", "/":
			i++
		default:
			fmt.Println("Wrong operator Please try again.")
			fmt.Print("Please enter operator :  ")
			fmt.Scan(&operator)
		}
	}
	return operator
}

func operandCheck(num string) (int64, error) {

	numChecked, err := strconv.ParseInt(num, 10, 64)

	return numChecked, err
}

