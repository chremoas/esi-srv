package handler

import (
	"fmt"
	"github.com/antihax/goesi"
	"github.com/chremoas/esi-srv/proto"
	"github.com/gregjones/httpcache"
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
)

type allianceServiceHandler struct {
	Client    client.Client
	ESIClient *goesi.APIClient
}

func NewAllianceServiceHandler() chremoas_esi.AllianceServiceHandler {
	httpClient := httpcache.NewMemoryCacheTransport().Client()
	return &allianceServiceHandler{ESIClient: goesi.NewAPIClient(httpClient, "chremoas-esi-srv Ramdar Chinken on TweetFleet Slack https://github.com/chremoas/esi-srv")}
}

func (eqh *allianceServiceHandler) GetAlliances(ctx context.Context, request *chremoas_esi.GetAlliancesRequest, response *chremoas_esi.GetAlliancesResponse) error {
	alliance, _, err := eqh.ESIClient.ESI.AllianceApi.GetAlliances(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("Had some kind of error getting the alliance '%s'\n", err)
	}

	response.Ids = alliance

	return nil
}

func (eqh *allianceServiceHandler) GetAllianceById(ctx context.Context, request *chremoas_esi.GetAllianceByIdRequest, response *chremoas_esi.GetAllianceByIdResponse) error {
	alliance, _, err := eqh.ESIClient.ESI.AllianceApi.GetAlliancesAllianceId(context.Background(), request.Id, nil)
	if err != nil {
		return fmt.Errorf("Had some kind of error getting the alliance '%s'\n", err)
	}

	response.Alliance = &chremoas_esi.Alliance{
		Name:         alliance.AllianceName,
		Ticker:       alliance.Ticker,
		ExecutorCorp: alliance.ExecutorCorp,
		DateFounded:  alliance.DateFounded.Unix(),
	}

	return nil
}

func (eqh *allianceServiceHandler) GetAllianceCorporationsById(ctx context.Context, request *chremoas_esi.GetAllianceCorporationsByIdRequest, response *chremoas_esi.GetAllianceCorporationsByIdResponse) error {
	alliance, _, err := eqh.ESIClient.ESI.AllianceApi.GetAlliancesAllianceIdCorporations(context.Background(), request.Id, nil)
	if err != nil {
		return fmt.Errorf("Had some kind of error getting the alliance '%s'\n", err)
	}

	response.Ids = alliance

	return nil
}

func (eqh *allianceServiceHandler) GetAllianceIconsById(ctx context.Context, request *chremoas_esi.GetAllianceIconsByIdRequest, response *chremoas_esi.GetAllianceIconsByIdResponse) error {
	alliance, _, err := eqh.ESIClient.ESI.AllianceApi.GetAlliancesAllianceIdIcons(context.Background(), request.Id, nil)
	if err != nil {
		return fmt.Errorf("Had some kind of error getting the alliance '%s'\n", err)
	}

	response.Icons = &chremoas_esi.Icons{
		PSixtyFour:   alliance.Px64x64,
		POneTwentyEight: alliance.Px128x128,
	}

	return nil
}

func (eqh *allianceServiceHandler) GetAllianceNamesByIds(ctx context.Context, request *chremoas_esi.GetAllianceNamesByIdsRequest, response *chremoas_esi.GetAllianceNamesByIdsResponse) error {
	alliance, _, err := eqh.ESIClient.ESI.AllianceApi.GetAlliancesNames(context.Background(), request.Ids, nil)
	if err != nil {
		return fmt.Errorf("Had some kind of error getting the alliance '%s'\n", err)
	}

	for _, v := range alliance {
		response.Names = append(response.Names, &chremoas_esi.AllianceNames{AllianceId: v.AllianceId, AllianceName: v.AllianceName})
	}

	return nil
}
