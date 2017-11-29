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

	chremoas_esi.RegisterAllianceQueryHandler(service.Server(), handler.NewAllianceQueryHandler())
	chremoas_esi.RegisterCorporationQueryHandler(service.Server(), handler.NewCorporationQueryHandler())
	chremoas_esi.RegisterCharacterQueryHandler(service.Server(), handler.NewCharacterQueryHandler())
	chremoas_esi.RegisterSearchQueryHandler(service.Server(), handler.NewSearchQueryHandler())

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
