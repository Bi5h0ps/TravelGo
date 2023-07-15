package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorResponse(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
		"data":  nil,
	})
}

func SuccessResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "Success",
		"data": data,
	})
}
