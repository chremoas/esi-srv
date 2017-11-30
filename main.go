package main

import (
	"fmt"
	"github.com/chremoas/services-common/config"
	"github.com/chremoas/esi-srv/proto"
	"github.com/chremoas/esi-srv/handler"
)

var Version = "1.0.0"

func main() {
	service := config.NewService(Version, "esi-srv", config.NilInit)

	chremoas_esi.RegisterAllianceServiceHandler(service.Server(), handler.NewAllianceQueryHandler())
	chremoas_esi.RegisterCorporationServiceHandler(service.Server(), handler.NewCorporationQueryHandler())
	chremoas_esi.RegisterCharacterServiceHandler(service.Server(), handler.NewCharacterQueryHandler())
	chremoas_esi.RegisterSearchServiceHandler(service.Server(), handler.NewSearchQueryHandler())

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
