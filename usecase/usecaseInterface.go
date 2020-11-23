package usecase

type UscaseInterface interface{
	Plus(float64,float64) float64
	Minus(float64,float64) float64
	Multi(float64,float64) float64
	Div(float64,float64) (float64,error)
}