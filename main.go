package main

import (
	"git.dev.tanikawa.com/go/go_user/handler"
	"git.dev.tanikawa.com/go/go_user/subscriber"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"

	example "git.dev.tanikawa.com/go/go_user/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.go_user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.go_user", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.go_user", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
