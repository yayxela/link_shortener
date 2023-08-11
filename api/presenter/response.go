package presenter

import (
	"github.com/gofiber/fiber/v2"
)

func Success(ctx *fiber.Ctx, data interface{}) error {
	response := &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
	return ctx.JSON(response)
}

func Fail(ctx *fiber.Ctx, err error, status int) error {
	response := &fiber.Map{
		"status": true,
		"data":   nil,
		"error":  err.Error(),
	}
	ctx.Status(status)
	return ctx.JSON(response)
}
