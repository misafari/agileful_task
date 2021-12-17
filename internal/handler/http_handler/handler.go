package http_handler

import (
	"agileful_task/internal/core/domain"
	"agileful_task/internal/core/port"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type queriesHttpHandler struct {
	queriesService port.QueriesService
}

func (q *queriesHttpHandler) GetQueriesCountHandler(c *fiber.Ctx) error {
	count, err := q.queriesService.GetCount()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": count,
	})
}

func (q *queriesHttpHandler) GetQueriesHandler(c *fiber.Ctx) error {
	v := new(domain.QueryOption)

	if err := c.QueryParser(v); err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	all, err := q.queriesService.GetAll(v)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"data": all,
	})
}

func NewQueriesHttpHandler(queriesService port.QueriesService) port.QueriesHttpHandler {
	return &queriesHttpHandler{
		queriesService: queriesService,
	}
}
