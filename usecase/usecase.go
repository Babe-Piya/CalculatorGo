package usecase

import "errors"

type CalUsecase struct{}

func CreateUsecase() UscaseInterface{
	return CalUsecase{}
}

func (uc CalUsecase) Plus(f float64, s float64) float64 {
	return f+s
}

func (uc CalUsecase) Minus(f float64, s float64) float64 {
	return f-s
}

func (uc CalUsecase) Multi(f float64, s float64) float64 {
	return f*s
}

func (uc CalUsecase) Div(f float64, s float64) (float64,error) {
	if s==0 {
		return 0,errors.New("can't div by zero")
	}
	return f/s,nil
}