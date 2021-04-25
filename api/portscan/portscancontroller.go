package portscan

import (
	"GOnsumer/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func GetPortScan(ctx *fiber.Ctx) error {
	var (
		service    = ctx.Locals("service").(*service.Service)
		ip         = ctx.Params("ip")
		port       = ctx.Params("port")
		portStatus bool
	)
	portStatus = service.PortChecker.Check(ip, port)
	if !portStatus {
		return ctx.SendString("closed")
	}
	return ctx.SendString("open")
}

func GetPortScanWS(c *websocket.Conn) {
	defer c.Close() //TODO timeout koyabilirsin kapatmak icin, bazen write icin error basmıyor, connection kapanmıyor.

	service := c.Locals("service").(*service.Service)
	service.Web.WS = true

	var msg []byte
	var err error

	for {
		msg = <-service.Transporter.KafkaTransporter
		if err = c.WriteMessage(1, msg); err != nil {
			return
		}
	}

}
