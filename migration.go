package main

import (
	"api_tinggal_nikah/db"
	"api_tinggal_nikah/migration"
	"api_tinggal_nikah/models"
	"fmt"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
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
	conn := db.GetDB()
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
		&migration.CreateUsersPackage{},
		&migration.CreateMusicMasterTable{},
		&migration.CreateMusicUserTable{},
		// Add other migration instances here if needed
	}
	for _, m := range migrations {

		if err := m.Up(conn); err != nil {
			panic("Migration failed: " + err.Error())
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("@Min12cibubur"), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	userAdmin := models.User{
		FullName: "admin",
		Email:    "admin@tinggalnikah.com",
		Password: string(hashedPassword),
		Role:     models.Admin,
	}

	if err := conn.Create(&userAdmin).Error; err != nil {
		fmt.Println(err.Error())
		return
	}

	// end migration
}
