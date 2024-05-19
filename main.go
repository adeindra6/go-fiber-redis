package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var cache = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("REDIS_ADDR"),
	Password: "", // Set redis password first
	DB:       0,
})

var ctx = context.Background()

func verifyCache(c *fiber.Ctx) error {
	id := c.Params("id")
	val, err := cache.Get(ctx, id).Bytes()
	if err != nil {
		return c.Next()
	}

	data := toJson(val)
	return c.JSON(fiber.Map{
		"Cached": data,
	})
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Can't load .env file")
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("It is working!")
	})

	app.Get("data/:id", verifyCache, func(c *fiber.Ctx) error {
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

		cacheErr := cache.Set(ctx, id, body, 10*time.Second).Err()
		if cacheErr != nil {
			return cacheErr
		}

		data := toJson(body)
		return c.JSON(fiber.Map{
			"Data": data,
		})
	})

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	app.Listen(port)
}
