package main

import (
	"log"

	"github.com/emicklei/go-restful"

	 saying "github.com/irvifa/go-micro-greeter/proto/saying"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-web"

	"context"
)

type Saying struct{}

var (
	cl saying.SayingService
)

func (s *Saying) Anything(req *restful.Request, rsp *restful.Response) {
	log.Print("Received Saying.Anything API request")
	rsp.WriteEntity(map[string]string{
		"message": "Hello, Stranger!",
	})
}

func (s *Saying) WithName(req *restful.Request, rsp *restful.Response) {
	log.Print("Received Saying.Hello API request")

	name := req.PathParameter("name")

	response, err := cl.Hello(context.TODO(), &saying.Request{
		Name: name,
	})

	if err != nil {
		rsp.WriteError(500, err)
	}

	rsp.WriteEntity(response)
}

func main() {
	// Create service
	service := web.NewService(
		web.Name("go.micro.api.greeter"),
	)

	service.Init()

	// setup Greeter Server Client
	cl = saying.NewSayingService("go-micro-greeter.irvi-test", client.DefaultClient)

	// Create RESTful handler
	say := new(Saying)
	ws := new(restful.WebService)
	wc := restful.NewContainer()
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)
	ws.Path("/greeter")
	ws.Route(ws.GET("/").To(say.Anything))
	ws.Route(ws.GET("/{name}").To(say.WithName))
	wc.Add(ws)

	// Register Handler
	service.Handle("/", wc)

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
