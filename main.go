package main

import (
	"GOnsumer/service"
	"fmt"
	"os/signal"
	"syscall"
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
	)
	if err != nil {
		panic(err)
	}

	signal.Notify(srv.Sigchan, syscall.SIGINT, syscall.SIGTERM)

	go srv.Kafka.Consume(func(msg []byte, _ string) {
		fmt.Println(string(msg))
	}, srv.Sigchan, srv.DoneCh)

	<-srv.DoneCh
}
