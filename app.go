package main

import (
	"api_tinggal_nikah/db"
	"api_tinggal_nikah/services/auth"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	_, err := db.GetDB()
	if err != nil {
		fmt.Println("Error initializing database:", err)
		return
	}
	defer db.CloseDB()

	e := echo.New()

	v1 := e.Group(os.Getenv("PREFIX_API_VERSION"))
	auth.Routes(v1)
	e.Logger.Fatal(e.Start(os.Getenv("APP_PORTS")))
}
