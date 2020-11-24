package cmdHandler

import (
	"fmt"
	"calculator/usecase"
)

type CmdHandler struct {
	uc usecase.UscaseInterface
}

func CreateCmdHandler(uc usecase.UscaseInterface) CmdHandler {
	return CmdHandler{uc: uc}
}
func FlushScanf() {
	var discard string
	fmt.Scanln(&discard)
}
func DataHandler(displayStr string) (float64, error) {
	var number float64
	fmt.Print(displayStr)
	_, err := fmt.Scanf("%f\n", &number)
	if err != nil {
		return 0, err
	}
	return number, nil
}
