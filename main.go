package main

import (
	"mygoim/Conf"
	"mygoim/DB"
	"mygoim/Routers"
)

func main() {
	Conf.LoadConfig()
	DB.InitMySQL()
	DB.InitRedis()
	Routers.InitRouters()
}
