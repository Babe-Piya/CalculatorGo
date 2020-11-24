package handler

import (
	"calculator/usecase"
	"strconv"
)
type Handler struct {
	uc usecase.UscaseInterface
}
func CreateCmdHandler(uc usecase.UscaseInterface) Handler {
	return Handler{uc: uc}
}

func CheckNum(s string) (float64, error){
	 num, err := strconv.ParseFloat(s, 64)
	 return num,err
}