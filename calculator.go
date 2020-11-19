package main

import (
	"fmt"
)

func main() {
	var numFirst int
	var operator string
	var numSecond int
	var result int
	fmt.Print("Please enter operator :  ")
	fmt.Scan(&operator)
	operator = operatorCheck(operator)
	fmt.Print("Please enter first number : ")
	fmt.Scan(&numFirst)
	numFirstType := fmt.Sprintf("%T", numFirst)
	operandCheck(numFirstType)
	fmt.Print("Please enter second number :  ")
	fmt.Scan(&numSecond)
	result = cal(operator, numFirst, numSecond)
	fmt.Printf("Result : %d", result)

}

func cal(operator string, numFirst int, numSecond int) int {
	var result int

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
	return result
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

func operandCheck(numFirstType string) int {
	for i := 0; i < 1; {
		if numFirstType != "int" {
			fmt.Println("Invalid number please try again.")
			fmt.Print("Please enter operator:  ")
			fmt.Scan(&numFirstType)
		} else {
			i++
		}
	}
	return num
}
