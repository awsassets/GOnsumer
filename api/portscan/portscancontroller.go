package portscan

import (
	"GOnsumer/internal/service"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func GetPortScan(ctx *fiber.Ctx) error {
	service := ctx.Locals("service").(*service.Service)
	service.Logger.Info("request ack.")
	return ctx.SendString("Hello, World 👋!")
}

func GetPortScanWS(c *websocket.Conn) {
	service := c.Locals("service").(*service.Service)
	service.Web.WS = true

	var err error
	for {
		msg := <-service.Transporter.KafkaTransporter
		if err = c.WriteMessage(1, msg); err != nil {
			log.Println("write:", err)
			break
		}
	}
}
