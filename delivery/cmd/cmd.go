package cmd

import (
	"fmt"
	"os"
	"calculator/constant"
	cmdHandler "calculator/delivery/cmd/handler"
	"calculator/usecase"
)


type Calculator struct {
	choice string
}

func Exec (){
	hasArg := false
	h := cmdHandler.CreateCmdHandler(usecase.CreateUsecase())
	 cal := Calculator{}
	if len(os.Args) > 1 {
		cal.choice = os.Args[1]

	}
	if len(os.Args) > 3 {
		hasArg = true
	}

	for {
		if cal.choice == "" {
			fmt.Print("enter operator : ")
			fmt.Scanf("%s\n", &cal.choice)
		}
		switch constant.ChoiseType(cal.choice) {

		case constant.PLUS:
			h.PlusHandler(hasArg)
		case constant.MINUS:
			h.MinusHandler(hasArg)
		case constant.MULTI:
			h.MultiHandler(hasArg)
		case constant.DIV:
			h.DivHandler(hasArg)
		default:
			cal.choice = ""
			hasArg = false
		}
		cal.choice = ""
		hasArg = false
	}
}