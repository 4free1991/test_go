package main

import (
	"app/lesson4/config"
	"app/lesson4/controller"
	"app/lesson4/pkg/logger"
	"app/lesson4/pkg/mongo"
	"app/lesson4/pkg/mq"
	"app/lesson4/pkg/server"
	"app/lesson4/pkg/shutdown"
	"github.com/gin-gonic/gin"
	"syscall"
)

func main() {

	logger.Init()

	config.InitConfig()

	mongo.InitMongoClient()

	mq.InitKafaka()

	var engine *gin.Engine
	server.Init(engine, func(engine *gin.Engine) {
		controller.Init(engine)
	})

	shutdown.Listen(syscall.SIGINT, syscall.SIGTERM)
}
