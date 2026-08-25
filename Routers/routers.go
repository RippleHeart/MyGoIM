package Routers

import (
	"github.com/gin-gonic/gin"
	"mygoim/Handlers"
)

func InitRouters() error {
	Engine := gin.Default()
	Engine.POST("/login", Handlers.Login)

}
