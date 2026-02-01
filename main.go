package main

import (
	"log"
	"os"
	"product-api/model"
	"product-api/utils/database"
	"strings"
	"time"

	"product-api/handler"
	"product-api/repository"
	"product-api/service"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}
	config := model.Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	app := fiber.New()
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":    "ok",
			"message":   "Server is running",
			"version":   "1.0.0",
			"timestamp": time.Now().Format(time.RFC3339),
		})
	})

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo, categoryRepo)
	productHandler := handler.NewProductHandler(productService)

	app.Get("/api/category", categoryHandler.GetAll)
	app.Get("/api/category/:id", categoryHandler.GetByID)
	app.Post("/api/category", categoryHandler.Create)
	app.Put("/api/category/:id", categoryHandler.Update)
	app.Delete("/api/category/:id", categoryHandler.Delete)

	app.Get("/api/product", productHandler.HandleProducts)
	app.Get("/api/product/:id", productHandler.GetByID)
	app.Post("/api/product", productHandler.Create)
	app.Put("/api/product/:id", productHandler.Update)
	app.Delete("/api/product/:id", productHandler.Delete)

	err = app.Listen(":" + config.Port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	log.Println("Server started on port", config.Port)
}
