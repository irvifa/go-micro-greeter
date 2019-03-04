package main

import (
	"fmt"

	"context"

	proto "github.com/irvifa/go-micro-greeter/proto/saying"
	micro "github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(micro.Name("client"))

	service.Init()

	greeter := proto.NewSayingService("greeter.irvi-test", service.Client())
	rsp, err := greeter.SayHello(context.TODO(), &proto.Request{Name: "John"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.Msg)

}
