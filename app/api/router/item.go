package router

import (
	"github.com/gofiber/fiber/v2"
	"hoangphuc.tech/hercules/app/api/dto"
	"hoangphuc.tech/hercules/app/api/handler"
	"hoangphuc.tech/hercules/app/middleware"
)

type ItemRouter struct{}

func (r ItemRouter) Register(root fiber.Router) {
	group := root.Group("/item")
	handle := handler.NewItemHandler()
	valid := middleware.NewValidator()

	// Special item
	group.Get("/id.:id", valid.Params(&dto.EntityID{}), handle.Get)
	group.Post("/id.:id", valid.Params(&dto.EntityID{}), valid.Body(&dto.ItemUpdated{}), handle.Update)
	group.Get("/:code", valid.Params(&dto.EntityCode{}), handle.GetByCode)

	// Create new
	group.Post("/", valid.Body(&dto.ItemCreated{}), handle.Create)

}