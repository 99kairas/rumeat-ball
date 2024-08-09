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
			Role:     "admin",
		},
	}

	for _, v := range admin {
		var exist models.User

		errCheck := DB.Where("email = ?", v.Email).First(&exist).Error

		if errCheck != nil {
			DB.Create(&v)
		}
	}

	// CATEGORY MENU SEEDERS
	categoryID1, _ := uuid.Parse("6eef10e0-d7e8-44eb-b7f3-21b9fac7ca1f")
	categoryID2, _ := uuid.Parse("1db38e8a-abc4-4b21-aa91-01df9bc32358")
	categoryID3, _ := uuid.Parse("d48463a3-c975-4691-8673-8c27ee36c765")
	categoryID4, _ := uuid.Parse("5ec4d7c7-6534-45a9-9443-b349ce29b840")
	categoryID5, _ := uuid.Parse("7513379d-27ae-4f7b-be21-d3d3d8d3ab58")
	categoryMenu := []models.Category{
		{
			ID:   categoryID1,
			Name: "Mie",
		},
		{
			ID:   categoryID2,
			Name: "Nasi Goreng",
		},
		{
			ID:   categoryID3,
			Name: "Ayam",
		},
		{
			ID:   categoryID4,
			Name: "Cumi",
		},
		{
			ID:   categoryID5,
			Name: "Minuman",
		},
	}

	for _, v := range categoryMenu {
		var exist models.Category

		errCheck := DB.Where("id = ?", v.ID).First(&exist).Error

		if errCheck != nil {
			DB.Create(&v)
		}
	}

	// MENU SEEDERS
	menu := []models.Menu{
		{
			ID:          uuid.MustParse("12e100bb-fa3a-416f-b5d9-fb13d1d659b5"),
			Name:        "Mie Bakso",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla egestas suscipit tincidunt. Sed at volutpat magna, vitae accumsan libero. Ut eget sapien sem. Vestibulum nec posuere massa, non iaculis erat. Vestibulum lobortis nisl quis tellus tristique semper. Suspendisse luctus velit purus, ac accumsan mi faucibus interdum. Mauris gravida nisi eget eros faucibus mattis. Proin lobortis consequat facilisis.",
			Image:       "https://res.cloudinary.com/dw3n2ondc/image/upload/v1721171625/Rumeat-Ball/i9cbktn0dibyacm1ooos.jpg",
			Price:       19000,
			Status:      "available",
			CategoryID:  categoryID1,
		},
		{
			ID:          uuid.MustParse("3f0d9bc9-aa78-4562-8a74-a6297d8f57fb"),
			Name:        "Mie Ayam",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla egestas suscipit tincidunt. Sed at volutpat magna, vitae accumsan libero. Ut eget sapien sem. Vestibulum nec posuere massa, non iaculis erat. Vestibulum lobortis nisl quis tellus tristique semper. Suspendisse luctus velit purus, ac accumsan mi faucibus interdum. Mauris gravida nisi eget eros faucibus mattis. Proin lobortis consequat facilisis.",
			Image:       "https://res.cloudinary.com/dw3n2ondc/image/upload/v1721171625/Rumeat-Ball/smvksc6yctg6for2rev0.jpg",
			Price:       23000,
			Status:      "available",
			CategoryID:  categoryID1,
		},
		{
			ID:          uuid.MustParse("3b48bef8-5967-4e95-852c-804490a2625d"),
			Name:        "Nasi Goreng Spesial",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla egestas suscipit tincidunt. Sed at volutpat magna, vitae accumsan libero. Ut eget sapien sem. Vestibulum nec posuere massa, non iaculis erat. Vestibulum lobortis nisl quis tellus tristique semper. Suspendisse luctus velit purus, ac accumsan mi faucibus interdum. Mauris gravida nisi eget eros faucibus mattis. Proin lobortis consequat facilisis.",
			Image:       "https://res.cloudinary.com/dw3n2ondc/image/upload/v1721171625/Rumeat-Ball/kgi1acvwafyshqfgxpp3.jpg",
			Price:       32000,
			Status:      "available",
			CategoryID:  categoryID2,
		},
		{
			ID:          uuid.MustParse("5a98e3d8-eef6-46ce-b25a-32c040f4c0a4"),
			Name:        "Nasi Goreng Seafood",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla egestas suscipit tincidunt. Sed at volutpat magna, vitae accumsan libero. Ut eget sapien sem. Vestibulum nec posuere massa, non iaculis erat. Vestibulum lobortis nisl quis tellus tristique semper. Suspendisse luctus velit purus, ac accumsan mi faucibus interdum. Mauris gravida nisi eget eros faucibus mattis. Proin lobortis consequat facilisis.",
			Image:       "https://res.cloudinary.com/dw3n2ondc/image/upload/v1721171625/Rumeat-Ball/nt2mekgib5jlhqrtfnre.jpg",
			Price:       34000,
			Status:      "available",
			CategoryID:  categoryID2,
		},
		{
			ID:          uuid.MustParse("a9332e8c-bbd7-454d-ad66-f9142cdc5241"),
			Name:        "Juice Buah Naga",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla egestas suscipit tincidunt. Sed at volutpat magna, vitae accumsan libero. Ut eget sapien sem. Vestibulum nec posuere massa, non iaculis erat. Vestibulum lobortis nisl quis tellus tristique semper. Suspendisse luctus velit purus, ac accumsan mi faucibus interdum. Mauris gravida nisi eget eros faucibus mattis. Proin lobortis consequat facilisis.",
			Image:       "https://res.cloudinary.com/dw3n2ondc/image/upload/v1721171625/Rumeat-Ball/fh5trk0djnyahkmwoyj7.jpg",
			Price:       15000,
			Status:      "available",
			CategoryID:  categoryID5,
		},
		{
			ID:          uuid.MustParse("929fb956-f8b5-4a86-91b2-fc8368c0a4d5"),
			Name:        "Es Cendol",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla egestas suscipit tincidunt. Sed at volutpat magna, vitae accumsan libero. Ut eget sapien sem. Vestibulum nec posuere massa, non iaculis erat. Vestibulum lobortis nisl quis tellus tristique semper. Suspendisse luctus velit purus, ac accumsan mi faucibus interdum. Mauris gravida nisi eget eros faucibus mattis. Proin lobortis consequat facilisis.",
			Image:       "https://res.cloudinary.com/dw3n2ondc/image/upload/v1721171625/Rumeat-Ball/bjpmjogwk2axtvdicmqe.jpg",
			Price:       12000,
			Status:      "available",
			CategoryID:  categoryID5,
		},
	}
	for _, v := range menu {
		var exist models.Category

		errCheck := DB.Where("id = ?", v.ID).First(&exist).Error

		if errCheck != nil {
			DB.Create(&v)
		}
	}
}
