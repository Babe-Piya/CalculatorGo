package cmdHandler

import (
	"fmt"
	"os"
	"strconv"
)

func (c *CmdHandler) DivHandler(hasArg bool) {
	var firstNumber, secondNumber float64
	var err error
	if !hasArg {
		firstNumber, err = DataHandler("enter first number : ")
		if err == nil {
			secondNumber, err = DataHandler("enter second number : ")
			if err != nil {
				FlushScanf()
				return
			}
		} else {
			FlushScanf()
			return
		}
	} else {
		firstNumber, _ = strconv.ParseFloat(os.Args[2], 64)
		secondNumber, _ = strconv.ParseFloat(os.Args[3], 64)
	}
	data, err := c.uc.Div(firstNumber, secondNumber)
	if err != nil {
		fmt.Println("err", err)
	} else {
		fmt.Println("result : ", data)
	}

}
