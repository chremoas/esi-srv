package main

import (
	"fmt"
	"github.com/abaeve/services-common/config"
	"github.com/chremoas/esi-srv/proto"
	"github.com/chremoas/esi-srv/handler"
)

var Version = "1.0.0"

func main() {
	service := config.NewService(Version, "esi-srv", nil)

	chremoas_esi.RegisterEntityQueryHandler(service.Server(), handler.NewEntityQueryHandler())

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
