package port

import "github.com/gofiber/fiber/v2"

type QueriesHttpHandler interface {
	GetQueriesCountHandler(*fiber.Ctx) error
	GetQueriesHandler(*fiber.Ctx) error
}