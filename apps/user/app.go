package main

import (
	"fmt"
	"os"
	"path"

	"github.com/Psuta90/api_tinggal_nikah/apps/user/routes"
	"github.com/Psuta90/api_tinggal_nikah/db"
	"github.com/Psuta90/api_tinggal_nikah/utils"
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

	e.Validator = utils.NewCustomValidator()

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	adapter, err := gormadapter.NewAdapterByDB(db.GetDB())
	if err != nil {
		panic(err)
	}

	enforcer, err := casbin.NewEnforcer(path.Join(cwd, "middleware", "casbin", "model.conf"), adapter)
	if err != nil {
		panic(err)
	}

	// uncomment this every you change column or add tables
	// migrations := []migration.Migration{
	// 	&migration.CreateUsersTable{},
	// 	&migration.CreateAcaraTable{},
	// 	&migration.CreateGalleryPhotosTable{},
	// 	&migration.CreateLoveStoryTable{},
	// 	&migration.CreateMempelaiPriaTable{},
	// 	&migration.CreateMempelaiWanitaTable{},
	// 	&migration.CreateCasbinRule{},
	// 	&migration.CreateGiftDigital{},
	// 	&migration.CreateGuestBook{},
	// 	&migration.CreatePackageCategory{},
	// 	&migration.CreatePackageTable{},
	// 	// Add other migration instances here if needed
	// }
	// for _, m := range migrations {
	// 	conn := db.GetDB()
	// 	if err := m.Up(conn); err != nil {
	// 		panic("Migration failed: " + err.Error())
	// 	}
	// }
	// end migration

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v1 := e.Group(os.Getenv("PREFIX_API_VERSION"))
	routes.Routes(v1, enforcer)
	e.Logger.Fatal(e.Start(os.Getenv("APP_PORTS")))
}
