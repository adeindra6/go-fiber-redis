package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Can't load .env file")
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("It is working!")
	})

	app.Get("data/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		res, err := http.Get("https://jsonplaceholder.typicode.com/users/" + id)
		if err != nil {
			return err
		}

		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		user := User{}
		parseErr := json.Unmarshal(body, &user)
		if parseErr != nil {
			return parseErr
		}

		return c.JSON(fiber.Map{
			"Data": user,
		})
	})

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	app.Listen(port)
}
