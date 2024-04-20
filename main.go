package main

import (
	"rumeat-ball/configs"
	"rumeat-ball/database"
	m "rumeat-ball/middlewares"
	"rumeat-ball/routes"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Local = loc
	configs.Init()
	database.InitDatabase()
	e := routes.New()
	//implement middleware logger
	m.LogMiddlewares(e)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
