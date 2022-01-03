package response

import (
	gin "github.com/gin-gonic/gin"
)

func JSON(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, data)
}