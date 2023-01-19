package location

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kwanok/spatial-query-study/api/utils"
)

type NearQuery struct {
	X  float64 `query:"x" validate:"required,number,min=-90,max=90"`
	Y  float64 `query:"y" validate:"required,number,min=-180,max=180"`
	Km float64 `query:"km" validate:"required,number,min=0,max=99999"`
}

func (q *NearQuery) Validate(c *fiber.Ctx) error {
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

func NearHandlerV1(c *fiber.Ctx) error {
	query := new(NearQuery)
	if err := query.Validate(c); err != nil {
		return err
	}

	return c.JSON(FetchNearLocationsV1(query))
}

func NearHandlerV2(c *fiber.Ctx) error {
	query := new(NearQuery)
	if err := query.Validate(c); err != nil {
		return err
	}

	return c.JSON(FetchNearLocationsV2(query))
}
