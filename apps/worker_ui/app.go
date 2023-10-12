package main

import (
	"api_tinggal_nikah/config"
	"api_tinggal_nikah/db"
	"api_tinggal_nikah/utils"

	"github.com/hibiken/asynqmon"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	utils.Client()
	db.InitRedisConnection()
}

func main() {
	e := echo.New()
	//
	redisConnection := utils.RedisClientOpt
	mon := asynqmon.New(asynqmon.Options{
		RootPath:     "/monitoring/tasks",
		RedisConnOpt: redisConnection,
	})

	config.ConfigCors(e)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/monitoring/tasks/*", echo.WrapHandler(mon))
	e.Logger.Fatal(e.Start(":3005"))
}
