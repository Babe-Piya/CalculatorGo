package rest


import (
	"github.com/gin-gonic/gin"
	cmdHandler "calculator/delivery/rest/handler"
	"calculator/usecase"
)
func Exec(){
	r := gin.Default()
	h := cmdHandler.CreateCmdHandler(usecase.CreateUsecase())

	r.GET("/", func(context *gin.Context){
		context.JSON(200,gin.H{
			"helloWorld":"Hi",
			"id":1,
		})
	})

	r.GET("/plus/:data1/:data2", h.PlusHandler)

	r.GET("/minus/:data1/:data2", h.MinusHandler)

	r.GET("/div/:data1/:data2", h.DivHandler)

	r.GET("/multi/:data1/:data2", h.MultiHandler)

	r.Run()
}