package main

import (
	"GOnsumer/api/portscan"
	"GOnsumer/internal/service"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

var (
	Build   = "unknown"
	Version = "unknown"
	Name    = "GOnsumer"
	srv     *service.Service
)

func main() {
	var err error

	srv, err = service.New(Name, Version,
		service.Kafka(),
		service.Logger(),
		service.PortChecker(),
		service.Web(
			map[string]func(c *fiber.Ctx) error{
				"/scan/:ip/:port": portscan.GetPortScan,
				"/ws":             websocket.New(portscan.GetPortScanWS),
			},
		),
	)
	if err != nil {
		panic(err)
	}

	signal.Notify(srv.Sigchan, syscall.SIGINT, syscall.SIGTERM)

	//Run consumer
	go srv.Kafka.Consume(func(msg []byte, topic string) {
		if srv.Web.WS {
			srv.Transporter.KafkaTransporter <- msg
		}
		srv.Logger.Info(fmt.Sprintf("Received msg from %s: %s", topic, string(msg)))
	}, srv.Sigchan, srv.DoneCh)

	//Run web server
	go srv.Web.Run(srv)

	<-srv.DoneCh
}
