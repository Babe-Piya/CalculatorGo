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
var read bool

func main() {

	if len(os.Args) > 3 {
		read = true
	}
	for {
		if read {
			if !operatorCheck(os.Args[1]) {
				read = false
				continue
			}
			numFirst, err := operandCheck(os.Args[2])
			if err != nil {
				fmt.Println("Invalid number Please try again.")
				read = false
				continue
			}

			numSecond, err := operandCheck(os.Args[3])
			if err != nil {
				fmt.Println("Invalid number Please try again.")
				read = false
				continue
			}

			cal(os.Args[1], numFirst, numSecond)
			read = false
		}

		fmt.Print("Please enter operator :  ")
		fmt.Scan(&operator)
		if !operatorCheck(operator) {
			continue
		}
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

func operatorCheck(operator string) bool {
	check := false
	switch operator {
	case "+", "-", "*", "/":
		check = true
	default:
		fmt.Println("Wrong operator Please try again.")
		check = false
	}
	return check
}

func operandCheck(num string) (int64, error) {

	numChecked, err := strconv.ParseInt(num, 10, 64)

	return numChecked, err
}
