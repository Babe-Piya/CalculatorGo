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

func cal(operator string, numFirst float64, numSecond float64) {
	switch operator {
	case "+":
		showResultPlus( numFirst,numSecond)
	case "-":
		showResultMinus( numFirst,numSecond)
	case "*":
		showResultMultiply( numFirst,numSecond)
	case "/":
		showResultDevided(numFirst,numSecond)
	}
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

func operandCheck(num string) (float64, error) {

	numChecked, err := strconv.ParseFloat(num, 64)

	return numChecked, err
}

func showResultDevided(numFirst float64, numSecond float64) {
	result := numFirst / numSecond
	if numSecond == 0 {
		fmt.Println("Can't devide by 0 ")
	} else {
	fmt.Printf("Result : %f\n", result)
	}
}

func showResultPlus(numFirst float64, numSecond float64) {
	result := numFirst + numSecond
	fmt.Printf("Result : %f\n", result)
	
}

func showResultMinus(numFirst float64, numSecond float64) {
	result := numFirst - numSecond
	fmt.Printf("Result : %f\n", result)
	
}

func showResultMultiply(numFirst float64, numSecond float64) {
	result := numFirst * numSecond
	fmt.Printf("Result : %f\n", result)
	
}
