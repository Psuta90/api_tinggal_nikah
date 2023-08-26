package main

import (
	"api_tinggal_nikah/db"
	"api_tinggal_nikah/modules/auth"
	"api_tinggal_nikah/utils"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db.InitDB()
	e.Validator = utils.NewCustomValidator()

	// uncomment this every you change column or add tables
	// migrations := []migration.Migration{
	// 	&migration.CreateUsersTable{},
	// 	&migration.CreateAcaraTable{},
	// 	&migration.CreateGalleryPhotosTable{},
	// 	&migration.CreateLoveStoryTable{},
	// 	&migration.CreateMempelaiPriaTable{},
	// 	&migration.CreateMempelaiWanitaTable{},
	// 	// Add other migration instances here if needed
	// }
	// for _, m := range migrations {
	// 	conn := db.GetDB()
	// 	if err := m.Up(conn); err != nil {
	// 		panic("Migration failed: " + err.Error())
	// 	}
	// }
	// end migration

	v1 := e.Group(os.Getenv("PREFIX_API_VERSION"))
	auth.Routes(v1)
	e.Logger.Fatal(e.Start(os.Getenv("APP_PORTS")))
}
