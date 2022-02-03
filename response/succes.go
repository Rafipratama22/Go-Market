package response

import (
	gin "github.com/gin-gonic/gin"
)

type Response interface {
	SuccessResponse(ctx *gin.Context, data interface{})
	CreateResponse(ctx *gin.Context, data interface{})
	BadRequestResponse(ctx *gin.Context, data interface{})
	InternalServerErrResponse(ctx *gin.Context, data interface{})
}

type responseSuccess struct {
	status int
	data interface{}
}


func SuccessResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, responseSuccess{status: 200, data: data})
}

func CreateResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(201, responseSuccess{status: 201, data: data})
}

func BadRequestResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(400, responseSuccess{status: 400, data: data})
}

func InternalServerErrResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(500, responseSuccess{status: 500, data: data})
}