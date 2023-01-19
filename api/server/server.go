package server

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kwanok/spatial-query-study/api/config"
	"github.com/kwanok/spatial-query-study/api/location"
)

type Server struct {
}

func Start() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	locations := app.Group("/locations")
	locations.Get("/v1/near", location.NearHandler)
	locations.Get("/v2/near", location.NearHandlerV2)
	locations.Get("/polygon", location.PolygonHandler)

	app.Listen(fmt.Sprintf(":%s", config.RuntimeConf.Server.Port))
}
