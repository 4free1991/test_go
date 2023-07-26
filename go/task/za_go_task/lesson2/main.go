package main

import (
	"app/lesson2/dao"
	"app/lesson2/router"
)

func main() {
	//初始化数据库
	err := dao.InitDb()
	if err != nil {
		return
	}

	r := router.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	err = r.Run(":8088")
	if err != nil {
		return
	}
}
