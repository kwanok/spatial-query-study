package server

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/kwanok/spatial-query-study/api/location"
	"github.com/spf13/viper"
)

type Server struct {
}

func Start() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	locations := app.Group("/locations")
	locations.Get("/v1/near", location.NearHandlerV1)
	locations.Get("/v2/near", location.NearHandlerV2)

	app.Listen(fmt.Sprintf(":%s", viper.Get("PORT")))
}
