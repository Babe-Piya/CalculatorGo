package cmdHandler

import (
	"fmt"
	"os"
	"strconv"
)

func (c *CmdHandler) PlusHandler(hasArg bool) {
	var firstNum, secondNum float64
	var err error
	if !hasArg {
		firstNum, err = DataHandler("enter first num : ")
		if err == nil {
			secondNum, err = DataHandler("enter second number : ")
			if err != nil {
				FlushScanf()
				return
			}

		} else {
			FlushScanf()
			return

		}
	} else {
		firstNum, _ = strconv.ParseFloat(os.Args[2], 64)
		secondNum, _ = strconv.ParseFloat(os.Args[3], 64)
	}
	fmt.Println("result : ",c.uc.Plus(firstNum, secondNum))
}
