package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-suriya/go-fiber-api/config"
	"github.com/go-suriya/go-fiber-api/database"
	"github.com/gofiber/fiber/v2"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("hello world 🌈")
}

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db := database.NewPostgresDatabase(*config)

	defer func() {
		if err := db.Close(); err != nil {
			fmt.Printf("Database shutdown error: %s\n", err)
		}
	}()

	// fiber instance
	app := fiber.New()

	// routes
	app.Get("/", helloWorld)

	go gracefulShutdown(app)

	app.Listen(config.GetPort())
}

func gracefulShutdown(app *fiber.App) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	d := time.Duration(5 * time.Second)
	fmt.Printf("Shutting down with a %s timeout...\n", d)

	ctx, cancel := context.WithTimeout(context.Background(), d)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		fmt.Printf("Fiber server shutdown error: %s\n", err)
	}

	fmt.Println("Gracefully shut down")
}
