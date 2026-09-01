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
	UserRG := Engine.Group("/users/:name")
	UserRG.Use(Handlers.VerifyJWT)
	{
		UserRG.GET("/friends", Handlers.GetFriends)
		UserRG.GET("/friends/:friendname")
		UserRG.POST("/friends/:friendname", Handlers.AddFriend)

		UserRG.POST("/groups/:groupname", Handlers.CreateGroup)
		UserRG.PATCH("/groups/:groupname", Handlers.EnterGroup)
		UserRG.GET("/groups/:groupname", Handlers.GetMembers)
		UserRG.GET("/hello", Handlers.Hello)
		UserRG.GET("/chat", Handlers.WSUpgrade)
	}
	err := Engine.Run("localhost:8080")
	if err != nil {
		log.Fatal("Engine.Run: ", err)
	}
}
