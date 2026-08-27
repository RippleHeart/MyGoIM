package Routers

import (
	"github.com/gin-gonic/gin"
	"log"
	"mygoim/Handlers"
)

func InitRouters() {
	Engine := gin.Default()
	Engine.POST("/login", Handlers.Login)
	Engine.POST("/register", Handlers.Register)
	UserRG := Engine.Group("/user/:name", Handlers.VerifyJWT)
	UserRG.Use()
	{
		UserRG.GET("/chat", Handlers.WSUpgrade)
		UserRG.GET("/hello", Handlers.Hello)
	}
	err := Engine.Run("localhost:8080")
	if err != nil {
		log.Fatal("Engine.Run: ", err)
	}
}
