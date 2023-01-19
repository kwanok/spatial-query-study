package location

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/kwanok/spatial-query-study/api/utils"
)

type NearQuery struct {
	X  float64 `query:"x" validate:"required,number,min=-180,max=180"`
	Y  float64 `query:"y" validate:"required,number,min=-90,max=90"`
	Km float64 `query:"km" validate:"required,number,min=0,max=99999"`
}

type PolygonQuery struct {
	X1 float64 `query:"x1" validate:"required,number,min=-180,max=180"`
	Y1 float64 `query:"y1" validate:"required,number,min=-90,max=90"`
	X2 float64 `query:"x2" validate:"required,number,min=-180,max=180"`
	Y2 float64 `query:"y2" validate:"required,number,min=-90,max=90"`
}

func (q *PolygonQuery) ConvertPolygon() string {
	return fmt.Sprintf("POLYGON((%f %f, %f %f, %f %f, %f %f, %f %f))", q.X1, q.Y1, q.X1, q.Y2, q.X2, q.Y2, q.X2, q.Y1, q.X1, q.Y1)
}

func Validate[T any](c *fiber.Ctx, q T) error {
	if err := c.QueryParser(q); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if errors := utils.ValidateStruct(q); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return nil
}

func NearHandler(c *fiber.Ctx) error {
	query := new(NearQuery)
	if err := Validate(c, query); err != nil {
		return err
	}

	return c.JSON(FetchNearLocationsV1(query))
}

func NearHandlerV2(c *fiber.Ctx) error {
	query := new(NearQuery)
	if err := Validate(c, query); err != nil {
		return err
	}

	return c.JSON(FetchNearLocationsV2(query))
}

func PolygonHandler(c *fiber.Ctx) error {
	query := new(PolygonQuery)
	if err := Validate(c, query); err != nil {
		return err
	}

	return c.JSON(FetchPolygonLocations(query))
}
