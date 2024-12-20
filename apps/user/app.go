package main

import (
	"api_tinggal_nikah/apps/user/routes"
	"api_tinggal_nikah/config"
	"api_tinggal_nikah/db"
	messagebroker "api_tinggal_nikah/message_broker"

	"api_tinggal_nikah/utils"
	"fmt"
	"os"
	"path"

	"github.com/casbin/casbin/v2"
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

	messagebroker.InitNATS()
	defer messagebroker.CloseNATS()

	e.Validator = utils.NewCustomValidator()

	_, err := config.InitializeMinioClient()
	if err != nil {
		panic(err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	adapter, err := gormadapter.NewAdapterByDB(db.GetDB())
	if err != nil {
		panic(err)
	}

	enforcer, err := casbin.NewEnforcer(path.Join(cwd, "../../", "middleware", "casbin", "model.conf"), adapter)
	if err != nil {
		panic(err)
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v1 := e.Group(os.Getenv("PREFIX_API_VERSION"))
	routes.CustomerRoutes(v1, enforcer)
	routes.AdminRoutes(v1, enforcer)
	routes.GlobalRoutes(v1)
	e.Logger.Fatal(e.Start(os.Getenv("APP_PORTS")))
}
