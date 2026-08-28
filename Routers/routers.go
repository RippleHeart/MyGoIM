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
		UserRG.GET("/friends", Handlers.GetFriends)
		UserRG.POST("/friends", Handlers.AddFriend)

		UserRG.PATCH("/group/:groupname", Handlers.CreateGroup)
		UserRG.POST("/group/:groupname", Handlers.EnterGroup)
		UserRG.GET("/group/:groupname", Handlers.GetMembers)
		UserRG.GET("/hello", Handlers.Hello)
		UserRG.GET("/chat", Handlers.WSUpgrade)
	}
	err := Engine.Run("localhost:8080")
	if err != nil {
		log.Fatal("Engine.Run: ", err)
	}
}
