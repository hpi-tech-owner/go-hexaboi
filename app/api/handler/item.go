package handler

import (
	"github.com/gofiber/fiber/v2"
	"hoangphuc.tech/go-hexaboi/app/api/dto"
	"hoangphuc.tech/go-hexaboi/infra/adapter"
	"hoangphuc.tech/go-hexaboi/infra/core"
)

type ItemHandler struct {
	repoItem adapter.ItemRepository
}

func NewItemHandler() *ItemHandler {
	return &ItemHandler{
		repoItem: adapter.ItemRepository{},
	}
}

func (h ItemHandler) GetByCode(c *fiber.Ctx) error {
	code := c.Params("code")
	item, err := h.repoItem.GetByCode(code)
	if err != nil {
		return err
	}
	return HJSON(c, item)
}

func (h ItemHandler) Get(c *fiber.Ctx) error {
	id, _ := core.Utils.ParseUint(c.Params("id"))
	item, err := h.repoItem.GetByID(id)
	if err != nil {
		return err
	}
	return HJSON(c, item)
}

func (h ItemHandler) Create(c *fiber.Ctx) error {
	// Parse payload as domain.Item
	d := new(dto.ItemCreated)
	if err := c.BodyParser(d); err != nil {
		return err
	}

	item := d.ToModel()

	// Create new item into repository
	err := h.repoItem.Create(item)
	if err != nil {
		return err
	}

	return HJSON(c, item)
}

func (h ItemHandler) Update(c *fiber.Ctx) error {
	id, _ := core.Utils.ParseUint(c.Params("id"))

	// Parse payload as domain.Item
	d := new(dto.ItemUpdated)
	if err := c.BodyParser(d); err != nil {
		return err
	}

	// Create new item into repository
	item := d.ToModel()
	err := h.repoItem.Update(id, item)
	if err != nil {
		return err
	}

	return HJSON(c, item)
}

func (h ItemHandler) Search(c *fiber.Ctx) error {
	var result []map[string]interface{}
	total, aggs, err := h.repoItem.Search(string(c.Body()), &result)
	if err != nil {
		return err
	}

	return HJSON(c, map[string]interface{}{
		"total":  total,
		"result": result,
		"aggs":   aggs,
	})
}

func (h ItemHandler) SearchIndex(c *fiber.Ctx) error {

	// Body can be a single element or a collection of elements
	var elements []interface{}
	if err := c.BodyParser(&elements); err != nil {
		var i interface{}
		if e := c.BodyParser(&i); e != nil {
			return e
		}
		elements = append(elements, i)
	}

	err := h.repoItem.SearchIndex(elements...)
	if err != nil {
		return err
	}
	return HJSON(c, nil)
}
