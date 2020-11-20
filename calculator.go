package main

import (
	"fmt"
	"strconv"
)

var first string
var second string
var operator string
var result int64

func main() {

	fmt.Print("Please enter operator :  ")
	fmt.Scan(&operator)
	operator = operatorCheck(operator)

	fmt.Print("Please enter first number : ")
	fmt.Scan(&first)
	numFirst := operandCheck(first)

	fmt.Print("Please enter second number :  ")
	fmt.Scan(&second)
	numSecond := operandCheck(second)

	result = cal(operator, numFirst, numSecond)
	fmt.Printf("Result : %d", result)

}

func cal(operator string, numFirst int64, numSecond int64) int64 {
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

func operandCheck(num string) int64 {
	var numChecked int64
	for i := 0; i < 1; {
		if isNumeric(num) {
			numChecked, _ = strconv.ParseInt(num, 10, 64)
			i++

		} else {
			fmt.Println("Wrong operator Please try again.")
			fmt.Print("Please enter operand :  ")
			fmt.Scan(&num)
		}
	}
	return numChecked
}

func isNumeric(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}
