package main

import (
	"web-go-skeleton/app/api/routes"
	"web-go-skeleton/app/api/components"
)


func main() {
	r := routes.InitRouter()
	//初始化工作
	components.InitLogger()
	components.InitRedis()

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
