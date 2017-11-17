package main

import (
	"fmt"
	"github.com/abaeve/services-common/config"
	"github.com/micro/go-micro"

	"github.com/abaeve/chremoas-command-template/command"
	proto "github.com/abaeve/chremoas/proto"
	"github.com/chremoas/esi-srv/proto"
	"github.com/chremoas/esi-srv/handler"
)

var Version string = "1.0.0"
var service micro.Service

func main() {
	service = config.NewService(Version, "esi", initialize)
	chremoas_esi.RegisterEntityQueryHandler(service.Server(), handler.NewEntityQueryHandler())

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
