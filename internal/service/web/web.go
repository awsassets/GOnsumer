package web

import (
	"github.com/gofiber/fiber/v2"
)

type (
	WebService struct {
		App    *fiber.App
		Routes []map[string]func(ctx *fiber.Ctx) error
	}
	Config struct {
	}
)

func (w *WebService) Run(srv interface{}) {
	w.App.Use(func(c *fiber.Ctx) error {
		c.Locals("service", srv)
		return c.Next()
	})
	for _, r := range w.Routes {
		for path, f := range r {
			w.App.Get(path, f)
		}
	}

	w.App.Listen(":1234")
}
