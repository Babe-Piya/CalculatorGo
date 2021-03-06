package handler
import (
	"github.com/gin-gonic/gin"
	"calculator/dto"
)


func (c *Handler) MinusHandler(context *gin.Context) {
	data1:= context.Param("data1")
	data2 := context.Param("data2")

	firstNum, err := CheckNum(data1)
	secondNum, err2 := CheckNum(data2)
	if err != nil {
		context.JSON(200,dto.Response{	
			Error :err.Error(),
		})	
		return
	}else if err2 != nil {
		context.JSON(200,dto.Response{	
			Error :err2.Error(),
		})	
		return
	}
	result := c.uc.Minus(firstNum, secondNum)
	 context.JSON(200,dto.Response{	
		Result :result,
	})		

}
