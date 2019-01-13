package main

import (
	"log"
	"time"

	saying "github.com/irvifa/go-micro-greeter/proto/saying"
	"github.com/micro/go-micro"

	"context"
)

type Saying struct{}

func (s *Saying) Hello(ctx context.Context, req *saying.Request, rsp *saying.Response) error {
	log.Print("Received Saying.Hello request")
	rsp.Msg = "Hello, " + req.Name + "!"
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.greeter"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	saying.RegisterSayingHandler(service.Server(), new(Saying))

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
