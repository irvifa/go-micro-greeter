package main

import (
	"fmt"

	"context"

	proto "github.com/irvifa/go-micro-greeter/proto/saying"
	micro "github.com/micro/go-micro"
)

/*
Example usage of top level service initialisation
*/

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("greeter.irvi-test"))

	service.Init()

	proto.RegisterSayingHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
