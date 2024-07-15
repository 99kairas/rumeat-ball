package main

import (
	"fmt"
	"log"
	"os"
	"rumeat-ball/configs"
	"rumeat-ball/database"
	m "rumeat-ball/middlewares"
	"rumeat-ball/repositories"
	"rumeat-ball/routes"
	"time"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Verify environment variables
	serverKey := os.Getenv("Server_Key")
	clientKey := os.Getenv("Client_Key")

	if serverKey == "" || clientKey == "" {
		log.Fatalf("Environment variables are not set correctly: Server_Key: %s, Client_Key: %s", serverKey, clientKey)
	}

	fmt.Println("Server Key:", serverKey)
	fmt.Println("Client Key:", clientKey)

	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Local = loc
	configs.Init()
	database.InitDatabase()
	e := routes.New()
	//implement middleware logger
	m.LogMiddlewares(e)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))

	c := cron.New()

	// Menambahkan tugas cron job yang berjalan setiap hari pukul 00:00
	c.AddFunc("@daily", func() {
		err := repositories.PermanentlyDeleteOldMenus(1 * 24 * time.Hour) // Menghapus data yang sudah lebih dari 1 hari
		if err != nil {
			log.Printf("Error deleting old menus: %v", err)
		} else {
			log.Println("Old menus deleted successfully")
		}
	})

	// Memulai scheduler
	c.Start()
	// Pastikan aplikasi tetap berjalan
	select {}
}
