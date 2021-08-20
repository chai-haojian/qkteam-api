package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"qkteam-api/logic"
	"qkteam-api/models"
)

func SubmitHandler(c *gin.Context) {
	//获取参数及参数校验
	p := new(models.Register)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Register with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//业务处理
	if err := logic.Submit(p); err != nil {
		zap.L().Error("logic.Submit(p) failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeError, "submit failed")
		return
	}
	//返回响应
	ResponseSuccess(c, nil)
}
