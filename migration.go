package main

import (
	"api_tinggal_nikah/db"
	"api_tinggal_nikah/migration"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	// cwd, err := os.Getwd()
	// if err != nil {
	// 	return
	// }

	err := godotenv.Load(".env.local")
	if err != nil {
		fmt.Println("Error loading custom_env.env file")
		return
	}

	db.InitDB()

	// uncomment this every you change column or add tables
	migrations := []migration.Migration{
		&migration.CreateUsersTable{},
		&migration.CreateAcaraTable{},
		&migration.CreateGalleryPhotosTable{},
		&migration.CreateLoveStoryTable{},
		&migration.CreateMempelaiPriaTable{},
		&migration.CreateMempelaiWanitaTable{},
		&migration.CreateCasbinRule{},
		&migration.CreateGiftDigital{},
		&migration.CreateGuestBook{},
		&migration.CreatePackageCategory{},
		&migration.CreatePackageTable{},
		&migration.CreateDomainTable{},
		&migration.TemplateMasterTable{},
		&migration.TemplateType{},
		&migration.TemplateUser{},
		&migration.CreateUsersTransactionTable{},
		// Add other migration instances here if needed
	}
	for _, m := range migrations {
		conn := db.GetDB()
		if err := m.Up(conn); err != nil {
			panic("Migration failed: " + err.Error())
		}
	}
	// end migration
}
