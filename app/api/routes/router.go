package routes

import (
	"github.com/gin-gonic/gin"
	"web-go-skeleton/app/api/controllers"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	//router.LoadHTMLGlob("views/*")
	//注册：
	router.GET("/",controllers.TestIndex)
	router.POST("/test/post",controllers.TestPost)
	router.POST("/test/log",controllers.TestLog)
	router.GET("/init",controllers.Migration)
	router.GET("/season/add",controllers.AddSeason)
	router.GET("/test/redis",controllers.TestRedis)

	return router

}
