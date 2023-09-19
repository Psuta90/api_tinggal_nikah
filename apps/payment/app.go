package main

import (
	"api_tinggal_nikah/apps/payment/routes"
	"api_tinggal_nikah/db"
	messagebroker "api_tinggal_nikah/message_broker"

	"api_tinggal_nikah/utils"
	"os"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type GormAdapterWrapper struct {
	*gormadapter.Adapter
}

func main() {
	e := echo.New()

	db.InitDB()
	db.InitRedisConnection()

	e.Validator = utils.NewCustomValidator()
	// nc, _ := nats.Connect("nats://localhost:6222")

	messagebroker.InitNATS()
	defer messagebroker.CloseNATS()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// go controller.MakePayment()

	v1 := e.Group(os.Getenv("PREFIX_API_VERSION"))
	routes.Routes(v1)
	e.Logger.Fatal(e.Start(os.Getenv("APP_PORTS")))
}
