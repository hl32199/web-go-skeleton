package controllers

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"web-go-skeleton/app/api/models"
	"github.com/araddon/dateparse"
	"web-go-skeleton/app/api/self_errors"
)


func AddSeason(ctx *gin.Context)  {
	startTime,_ := dateparse.ParseLocal("2019-09-01")
	endTime,_ := dateparse.ParseLocal("2019-09-30")
	season := models.Season{Month:201909,StartTime:startTime,EndTime:endTime}
	res := models.DB.Save(&season)

	if models.DuplicateEntry(res.Error) {
		ReturnError(ctx,self_errors.ErrSeasonAlreadyExist)
		return
	}

	if res.Error != nil {
		fmt.Println(res.Error)
		ReturnError(ctx,self_errors.ErrServiceNotAvailable)
		return
	}

	ReturnSuccess(ctx)
	return
}

