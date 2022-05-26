package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"hoangphuc.tech/hercules/infra/core"
)

// Response anything
func HJSON(c *fiber.Ctx, data interface{}) error {
	if hpiResult, ok := data.(core.HPIResult); ok {
		return hJSON(c, hpiResult)
	}
	status := c.Response().StatusCode()
	return hJSON(c, core.HPIResult{
		Status:  status,
		Data:    data,
		Message: utils.StatusMessage(status),
	})
}

func hJSON(c *fiber.Ctx, data core.HPIResult) error {
	if data.Status <= 0 {
		data.Status = c.Response().StatusCode()
	}

	if data.Message == "" {
		data.Message = utils.StatusMessage(data.Status)
	}

	return c.JSON(data)
}

// Centrialize errors
func HError(c *fiber.Ctx, status int, errCode string, err error, data ...interface{}) error {
	// Status code defaults to 500
	code := status
	if code < 400 {
		code = c.Response().StatusCode()
		if code < 400 {
			code = fiber.StatusInternalServerError
		}
	}

	// Retrieve the custom status code if it's an fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	// Retrieve the custom status code if it's an *core.HPIResult
	if e, ok := err.(*core.HPIResult); ok {
		code = e.Status
	}

	result := core.HPIResult{
		Status:    code,
		Message:   err.Error(),
		ErrorCode: errCode,
	}

	if len(data) > 0 {
		result.Data = data[0]
	}

	return HJSON(c, result)
}