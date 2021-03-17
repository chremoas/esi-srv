package main

import (
	"fmt"

	"github.com/chremoas/services-common/config"
	"go.uber.org/zap"

	chremoasPrometheus "github.com/chremoas/services-common/prometheus"

	"github.com/chremoas/esi-srv/handler"
	"github.com/chremoas/esi-srv/proto"
)

var (
	Version = "SET ME YOU KNOB"
	name    = "esi"
	logger  *zap.Logger
)

func main() {
	var err error

	// TODO pick stuff up from the config
	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	logger.Info("Initialized logger")

	go chremoasPrometheus.PrometheusExporter(logger)

	service := config.NewService(Version, "srv", name, config.NilInit)

	chremoas_esi.RegisterAllianceServiceHandler(service.Server(), handler.NewAllianceServiceHandler())
	chremoas_esi.RegisterCorporationServiceHandler(service.Server(), handler.NewCorporationServiceHandler())
	chremoas_esi.RegisterCharacterServiceHandler(service.Server(), handler.NewCharacterServiceHandler())
	chremoas_esi.RegisterSearchServiceHandler(service.Server(), handler.NewSearchServiceHandler())

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
