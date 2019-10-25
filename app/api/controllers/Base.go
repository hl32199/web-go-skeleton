package controllers

import (
	"github.com/gin-gonic/gin"
	"web-go-skeleton/app/api/self_errors"
	"net/http"
	"fmt"
)

//type BaseController struct {
//
//}

const (
	successCode = 0
	successMessage = "success"
	defaultErrMsg = "未知错误"
)

type ResultMsg struct {
	Code int64         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

/**
处理成功，响应数据
 */
func ReturnSuccess(ctx *gin.Context,data...interface{})  {
	res := ResultMsg{
		Code:successCode,
		Msg:successMessage,
	}

	if len(data) > 0 {
		res.Data = data[0]
	}

	ctx.JSON(http.StatusOK,res)
}

/*
响应错误信息
 */
func ReturnError(ctx *gin.Context,errCode int64)  {
	errMsg := ResultMsg{Code:errCode,Msg:defaultErrMsg}
	if msg,ok := self_errors.Messages[errCode];ok {
		errMsg.Msg = msg
	}

	ctx.JSON(http.StatusOK,errMsg)
}

/*
验证参数
 */
func ValidateParams(ctx *gin.Context,paramsObj interface{}) bool{
	err := ctx.ShouldBind(paramsObj)
	if err != nil{
		fmt.Println(err.Error())
		ReturnError(ctx,self_errors.ErrInvalidParams)
		return false
	}
	return true
}

