package handler

import (
	"fmt"
	"github.com/antihax/goesi"
	"github.com/chremoas/esi-srv/proto"
	"github.com/gregjones/httpcache"
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
)

type corporationServiceHandler struct {
	Client    client.Client
	ESIClient *goesi.APIClient
}

func NewCorporationServiceHandler() chremoas_esi.CorporationServiceHandler {
	httpClient := httpcache.NewMemoryCacheTransport().Client()
	return &corporationServiceHandler{ESIClient: goesi.NewAPIClient(httpClient, "chremoas-esi-srv Ramdar Chinken on TweetFleet Slack https://github.com/chremoas/esi-srv")}
}

func (eqh *corporationServiceHandler) GetCorporationById(ctx context.Context, request *chremoas_esi.GetCorporationByIdRequest, response *chremoas_esi.GetCorporationByIdResponse) error {
	corporation, _, err := eqh.ESIClient.ESI.CorporationApi.GetCorporationsCorporationId(context.Background(), request.Id, nil)
	if err != nil {
		return fmt.Errorf("Had some kind of error getting the corporation '%s'\n", err)
	}

	response.Corporation = &chremoas_esi.Corporation{
		Name:         corporation.Name,
		Ticker:       corporation.Ticker,
		MemberCount:  corporation.MemberCount,
		CeoId:        corporation.CeoId,
		AllianceId:   corporation.AllianceId,
		Description:  corporation.Description,
		TaxRate:      corporation.TaxRate,
		CreationDate: corporation.DateFounded.Unix(),
		CreatorId:    corporation.CreatorId,
		Url:          corporation.Url,
		FactionId:      corporation.FactionId,
	}

	return nil
}
