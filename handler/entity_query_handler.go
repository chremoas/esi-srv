package handler

import (
	"fmt"
	"github.com/antihax/goesi"
	"github.com/chremoas/esi-srv/proto"
	"github.com/gregjones/httpcache"
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
)

type entityQueryHandler struct {
	Client    client.Client
	ESIClient *goesi.APIClient
}

func NewEntityQueryHandler() chremoas_esi.EntityQueryHandler {
	httpClient := httpcache.NewMemoryCacheTransport().Client()
	return &entityQueryHandler{ESIClient: goesi.NewAPIClient(httpClient, "chremoas-esi-srv Ramdar Chinken on TweetFleet Slack https://github.com/chremoas/esi-srv")}
}

func (eqh *entityQueryHandler) GetAlliance(ctx context.Context, request *chremoas_esi.EntityQueryRequest, response *chremoas_esi.AllianceResponse) error {
	alliance, _, err := eqh.ESIClient.ESI.AllianceApi.GetAlliancesAllianceId(context.Background(), request.EntityId, nil)
	if err != nil {
		return fmt.Errorf("Had some kind of error getting the alliance '%s'\n", err)
	}

	response = &chremoas_esi.AllianceResponse{
		Id: request.EntityId,
		Name: alliance.AllianceName,
		Ticker: alliance.Ticker,
		ExecutorCorp: alliance.ExecutorCorp,
		DateFounded: alliance.DateFounded.Unix(),
	}

	return nil
}

func (eqh *entityQueryHandler) GetCorporation(ctx context.Context, request *chremoas_esi.EntityQueryRequest, response *chremoas_esi.CorporationResponse) error {
	corporation, _, err := eqh.ESIClient.ESI.CorporationApi.GetCorporationsCorporationId(context.Background(), request.EntityId, nil)
	if err != nil {
		return fmt.Errorf("Had some kind of error getting the corporation '%s'\n", err)
	}

	response = &chremoas_esi.CorporationResponse{
		Id: request.EntityId,
		Name: corporation.CorporationName,
		Ticker: corporation.Ticker,
		MemberCount: corporation.MemberCount,
		CeoId: corporation.CeoId,
		AllianceId: corporation.AllianceId,
		Description: corporation.CorporationDescription,
		TaxRate: corporation.TaxRate,
		CreationDate: corporation.CreationDate.Unix(),
		CreatorId: corporation.CreatorId,
		Url: corporation.Url,
		Faction: corporation.Faction,
	}

	return nil
}

func (eqh *entityQueryHandler) GetCharacter(ctx context.Context, request *chremoas_esi.EntityQueryRequest, response *chremoas_esi.CharacterResponse) error {
	character, _, err := eqh.ESIClient.ESI.CharacterApi.GetCharactersCharacterId(context.Background(), request.EntityId, nil)
	if err != nil {
		return fmt.Errorf("Had some kind of error getting the character '%s'\n", err)
	}

	response = &chremoas_esi.CharacterResponse{
		Id: request.EntityId,
		Name: character.Name,
		Description: character.Description,
		CorporationId: character.CorporationId,
		AllianceId: character.AllianceId,
		Birthday: character.Birthday.Unix(),
		Gender: character.Gender,
		RaceId: character.RaceId,
		BloodlineId: character.BloodlineId,
		AncestryId: character.AncestryId,
		SecurityStatus: character.SecurityStatus,
		// His docs say this exists, but it doesn't appear to
		//FactionId: character.FacrionId,
	}

	return nil
}

func (eqh *entityQueryHandler) GetSearch(ctx context.Context, request *chremoas_esi.EntitySearchRequest, response *chremoas_esi.SearchResponse) error {

	return nil
}
