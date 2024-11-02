package main

import (
	"github.com/go-suriya/go-fiber-api/database"
	"github.com/gofiber/fiber/v2"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("hello world 🌈")
}

func getFoods(c *fiber.Ctx) error {
	return c.JSON(foods)
}

func getFoodByID(c *fiber.Ctx) error {
	id := c.Params("id")
	for _, food := range foods {
		if food.ID == id {
			return c.JSON(food)
		}
	}
	return nil
}

func createFood(c *fiber.Ctx) error {
	var food Food
	if err := c.BodyParser(&food); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	foods = append(foods, food)
	return c.JSON(food)
}

func updateFoodByID(c *fiber.Ctx) error {
	var editFood EditFood
	if err := c.BodyParser(&editFood); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	id := c.Params("id")
	for i := range foods {
		if foods[i].ID == id {
			foods[i].Name = editFood.Name
			foods[i].Price = editFood.Price
			return c.JSON(foods[i])
		}
	}
	return nil
}

func deleteFoodByID(c *fiber.Ctx) error {
	id := c.Params("id")
	for i, food := range foods {
		if food.ID == id {
			foods = append(foods[:i], foods[i+1:]...)
			return c.SendString("delete success")
		}
	}
	return nil
}

func main() {
	database.ConnectDB()

	// fiber instance
	app := fiber.New()

	// routes
	app.Get("/", helloWorld)
	app.Get("/foods", getFoods)
	app.Get("/foods/:id", getFoodByID)
	app.Post("/foods", createFood)
	app.Put("/foods/:id", updateFoodByID)
	app.Delete("/foods/:id", deleteFoodByID)

	app.Listen("0.0.0.0:3000")
}

type Food struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price uint   `json:"price"`
}
type EditFood struct {
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

var foods = []Food{
	{ID: "1", Name: "ต้มยำกุ้ง", Price: 140},
	{ID: "2", Name: "ไก่ทอด", Price: 100},
	{ID: "3", Name: "ก๋วยเตี๋ยว", Price: 30},
	{ID: "4", Name: "เบอร์เกอร์", Price: 149},
}
