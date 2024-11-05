package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-suriya/go-fiber-api/database"
	"github.com/gofiber/fiber/v2"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("hello world 🌈")
}

func main() {
	database.ConnectDB()
	defer database.CloseDB()

	// fiber instance
	app := fiber.New()

	// routes
	app.Get("/", helloWorld)

	go func() {
		if err := app.Listen("0.0.0.0:" + "3000"); err != nil {
			fmt.Printf("Fiber server Listen error: %s\n", err)
		}
	}()

	gracefulShutdown(app)
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
