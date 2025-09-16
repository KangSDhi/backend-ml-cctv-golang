package main

import (
	"backend-ml-cctv-golang/config"
	"backend-ml-cctv-golang/entity"
	"backend-ml-cctv-golang/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("Successfully loaded .env file")
}

func loadDatabase() {
	config.InitDB()
	err := config.DB.AutoMigrate(&entity.CCTV{})
	if err != nil {
		log.Fatal("Error Migrate CCTV", err)
	}
	log.Println("Successfully Migrate CCTV")
}

func serveApplication() {

	app := fiber.New()

	apiRoutes := app.Group("/api/v1")

	apiRoutes.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"http_status": fiber.StatusOK,
			"message":     "OK",
		})
	})

	routes.SetupCCTVRouter(apiRoutes)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal("Server Error", err)
	}
}
