package main

import (
	"api_tinggal_nikah/config"
	"api_tinggal_nikah/db"
	"api_tinggal_nikah/utils"

	"github.com/hibiken/asynq"
	"github.com/hibiken/asynqmon"
	"github.com/labstack/echo/v4"
)

func init() {
	utils.Client()
	db.InitRedisConnection()
}

func main() {
	e := echo.New()
	//

	mon := asynqmon.New(asynqmon.Options{
		RootPath:     "/monitoring/tasks",
		RedisConnOpt: asynq.RedisClientOpt{Addr: "redis:6379", DB: 1},
	})

	config.ConfigCors(e)

	e.Any("/monitoring/tasks/*", echo.WrapHandler(mon))
	e.Logger.Fatal(e.Start(":3005"))
}
