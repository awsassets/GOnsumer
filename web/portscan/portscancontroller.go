package portscan

import (
	"GOnsumer/internal/service"

	"github.com/gofiber/fiber/v2"
)

func GetPortScan(ctx *fiber.Ctx) error {
	service := ctx.Locals("service").(*service.Service)
	service.Logger.Info("request ack.")
	return ctx.SendString("Hello, World 👋!")
}
