package main

import (
	"fmt"
	"github.com/chremoas/services-common/config"
	"github.com/chremoas/esi-srv/proto"
	"github.com/chremoas/esi-srv/handler"
)

var Version = "SET ME YOU KNOB"
var name = "esi"

func main() {
	service := config.NewService(Version, "srv", name, config.NilInit)

	chremoas_esi.RegisterAllianceServiceHandler(service.Server(), handler.NewAllianceServiceHandler())
	chremoas_esi.RegisterCorporationServiceHandler(service.Server(), handler.NewCorporationServiceHandler())
	chremoas_esi.RegisterCharacterServiceHandler(service.Server(), handler.NewCharacterServiceHandler())
	chremoas_esi.RegisterSearchServiceHandler(service.Server(), handler.NewSearchServiceHandler())

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
