package main

import (
	"fmt"
	"github.com/abaeve/services-common/config"
	"github.com/micro/go-micro"

	proto "github.com/abaeve/chremoas/proto"
	"github.com/abaeve/chremoas-command-template/command"
)

var Version string = "1.0.0"
var service micro.Service

func main ( ) {
	service = config.NewService(Version, "esi", initialize)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

// This function is a callback from the config.NewService function.  Read those docs
func initialize(config *config.Configuration) error {
	proto.RegisterCommandHandler(service.Server(),
		command.NewCommand(config.Name),
	)

	return nil
}