package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"time"
	"web-go-skeleton/app/api/components"
	"web-go-skeleton/app/api/models"
	"web-go-skeleton/app/api/self_errors"
	"web-go-skeleton/library/glog"
)

type Params struct {
	Name  string `form:"name" json:"name" binding:"required"`
	Id    uint   `form:"id" json:"id" binding:"required,min=5"`
	Items []Item `form:"items" json:"items" binding:"required,dive"` //dive用于让slice、array、map、嵌套的struct深入下一层
}

type Item struct {
	ProductId uint `form:"product_id" json:"product_id" binding:"required"`
	Number    uint `form:"number" json:"number" binding:"required"`
}

func TestIndex(ctx *gin.Context) {
	var seasons []models.Season
	res := models.DB.Find(&seasons)

	if res.Error != nil {
		ReturnError(ctx, self_errors.ErrServiceNotAvailable)
		return
	}

	ReturnSuccess(ctx, seasons)
	return
}

func TestPost(ctx *gin.Context) {
	var params Params
	if !ValidateParams(ctx, &params) {
		return
	}

	fmt.Println(params.Name)
	fmt.Println(params.Id)

	ReturnSuccess(ctx)
	return
}

func TestLog(ctx *gin.Context) {
	fmt.Fprintln(gin.DefaultWriter, "aaaaaaaabbbb")
	log.Println("log.Println")
	glog.Error("test glog error")
	glog.Info("test glog info")
	return
}

func TestRedis(ctx *gin.Context) {
	//err := components.Redis.Set("hello", "world", 0).Err()
	//if err != nil {
	//	glog.Error(err)
	//	ReturnError(ctx,self_errors.ErrServiceNotAvilable)
	//}
	var ch = make(chan string, 10)
	var list = make([]string, 0, 100)

	for i := 0; i < 100; i++ {
		go func(i int) {
			randNum := rand.Intn(10)
			time.Sleep(time.Duration(randNum) * time.Second)
			val, err := components.Redis.Get("hello").Result()
			if err != nil {
				glog.Error(err)
			}

			glog.Infof("%d-%d-%s", i, randNum, val)
			ch <- val
		}(i)
	}

	for i := 0; i < 100; i++ {
		list = append(list, <-ch)
	}

	ReturnSuccess(ctx)
	return
}

func Migration(ctx *gin.Context) {
	res := models.DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&models.Season{})
	if res.Error != nil {
		ctx.JSON(200, gin.H{
			"message": res.Error.Error(),
		})
	}

	ctx.String(200, "create table success")
	return
}
