package main

import (
	"fmt"
	"inibackend/config"
	"inibackend/router"
	"os"
	"strings"

	_ "inibackend/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func init() {
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			fmt.Println("Error loading .env file")
		}
	} else {
		fmt.Println(".env file not found, using default environment variables")
	}
}
// @title SWAGGER PEMROGRAMAN III
// @version 1.0
// @description Swagger Untuk Backend menggunakan golang dan framework Fiber

// @contact.name API Support
// @contact.url https://github.com/rrq-dev
// @contact.email raihanaditya1506@gmail

// @host https://backend-pemrogiii.onrender.com:6969
// @BasePath /
// @schemes http https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Join(config.GetAllowedOrigins(), ","),
		AllowHeaders:    "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
	}))

	router.SetupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "Not Found",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	fmt.Println(app.Listen(":" + port))
}