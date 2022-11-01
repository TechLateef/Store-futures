package controllers

import (
	"net/http"
	"strconv"

	"github.com/adewumi0550/pro_work/m/v2/dto"
	"github.com/adewumi0550/pro_work/m/v2/helper"
	"github.com/adewumi0550/pro_work/m/v2/services"
	"github.com/gin-gonic/gin"
)

type StoreController interface {
	UpdateStore(c *gin.Context)
	Enable(c *gin.Context)
	Disable(c *gin.Context)

}

type storeController struct {
	storeService services.Storeservice
	jwtService   services.JWTService
}

func NewStoreController(storeServ services.Storeservice, jwtServ services.JWTService) StoreController {
	return &storeController{
		storeService: storeServ,
		jwtService:   jwtServ,
	}
}


func (c *storeController) UpdateStore(ctx *gin.Context){
	id, err :=strconv.ParseUint(ctx.Param("ID"), 0,0)
	if err != nil {
		res := helper.BuildErrorResponse("NO Store found", err.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return

		}
		var store entities.store = c.storeService.UpdateStore(id)
		if (store != entities.store{}) {
			res := helper.BuildErrorResponse("Data not found", "No data with such id")
			ctx.JSON(http.StatusNotFound, res)
			
		} else{
			res := helper.BuildResponse(true, "Ok", store)
			ctx.JSON(http.StatusOK, res)
		}

	}



func (c *storeController)


func (c *storeController) UpdateStore(context *gin.Context) {
	var storeUpdateDTO dto.StoreUpdateDTO
	errDTO := context.ShouldBind(&storeUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

// func NewBookController(bookServ service.BookService, jwtServ service.JWTService) BookController {
// 	return &bookController{
// 		bookService: bookServ,
// 		jwtService:  jwtServ,
// 	}
// }
