package web

import (
	"github.com/gofiber/fiber/v2"
)

type (
	WebService struct {
		App    *fiber.App
		Routes []map[string]func(ctx *fiber.Ctx) error
		WS     bool
	}
	Config struct {
	}
)

func (w *WebService) Run(srv interface{}) {
	//init service
	w.App.Use(func(c *fiber.Ctx) error {
		c.Locals("service", srv)
		return c.Next()
	})

	w.App.Use("/ws", func(ctx *fiber.Ctx) error {
		ctx.Locals("allowed", true)
		return ctx.Next()
	})

	//init routes
	for _, r := range w.Routes {
		for path, f := range r {
			w.App.Get(path, f)
		}
	}
	w.App.Static("*", "../../web")
	w.App.Listen(":1234")
}
