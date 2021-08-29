package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"qkteam-api/logic"
)

func GetAllPostHandler(c *gin.Context) {
	data, err := logic.GetAllPost()
	if err != nil {
		zap.L().Error("logic.GetAllPost failed", zap.Error(err))
		fmt.Println(err)
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
