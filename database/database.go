package database

import (
	"fmt"
	"rumeat-ball/configs"
	"rumeat-ball/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDatabase() {
	InitDB()
	InitialMigration()
	Seeders()
}

type DbSetup struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() {
	database := DbSetup{
		DB_Username: configs.DB_USERNAME,
		DB_Password: configs.DB_PASSWORD,
		DB_Port:     configs.DB_PORT,
		DB_Host:     configs.DB_HOST,
		DB_Name:     configs.DB_NAME,
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC",
		database.DB_Username,
		database.DB_Password,
		database.DB_Host,
		database.DB_Port,
		database.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(
		&models.User{},
		&models.Menu{},
		&models.Order{},
		&models.Transaction{},
		&models.Rating{},
		&models.DetailOrder{},
	)
}

func Seeders() {
	// ADMIN SEEDERS
	adminPasswordHash, err := bcrypt.GenerateFromPassword([]byte("Admin@123"), bcrypt.DefaultCost)

	if err != nil {
		return
	}
	admin := []models.User{
		{
			ID:       uuid.New(),
			Email:    "admin1@admin.com",
			Password: string(adminPasswordHash),
			Name:     "Admin 1",
			Address:  "Jl. Jendral Sudirman",
			Phone:    "08123456789",
			Status:   "verified",
		},
	}

	for _, v := range admin {
		var exist models.User

		errCheck := DB.Where("email = ?", v.Email).First(&exist).Error

		if errCheck != nil {
			DB.Create(&v)
		}
	}
}
