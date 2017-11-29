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

func NewAllianceQueryHandler() chremoas_esi.AllianceQueryHandler {
	httpClient := httpcache.NewMemoryCacheTransport().Client()
	return &entityQueryHandler{ESIClient: goesi.NewAPIClient(httpClient, "chremoas-esi-srv Ramdar Chinken on TweetFleet Slack https://github.com/chremoas/esi-srv")}
}

func NewCorporationQueryHandler() chremoas_esi.CorporationQueryHandler {
	httpClient := httpcache.NewMemoryCacheTransport().Client()
	return &entityQueryHandler{ESIClient: goesi.NewAPIClient(httpClient, "chremoas-esi-srv Ramdar Chinken on TweetFleet Slack https://github.com/chremoas/esi-srv")}
}

func NewCharacterQueryHandler() chremoas_esi.CharacterQueryHandler {
	httpClient := httpcache.NewMemoryCacheTransport().Client()
	return &entityQueryHandler{ESIClient: goesi.NewAPIClient(httpClient, "chremoas-esi-srv Ramdar Chinken on TweetFleet Slack https://github.com/chremoas/esi-srv")}
}

func NewSearchQueryHandler() chremoas_esi.SearchQueryHandler {
	httpClient := httpcache.NewMemoryCacheTransport().Client()
	return &entityQueryHandler{ESIClient: goesi.NewAPIClient(httpClient, "chremoas-esi-srv Ramdar Chinken on TweetFleet Slack https://github.com/chremoas/esi-srv")}
}

func (eqh *entityQueryHandler) GetAlliances(ctx context.Context, request *chremoas_esi.EmptyRequest, response *chremoas_esi.AlliancesResponse) error {
	alliance, _, err := eqh.ESIClient.ESI.AllianceApi.GetAlliances(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("Had some kind of error getting the alliance '%s'\n", err)
	}

	response.AllianceIds = alliance

	return nil
}

func (eqh *entityQueryHandler) GetAllianceId(ctx context.Context, request *chremoas_esi.AllianceRequest, response *chremoas_esi.AllianceIdResponse) error {
	alliance, _, err := eqh.ESIClient.ESI.AllianceApi.GetAlliancesAllianceId(context.Background(), request.AllianceId, nil)
	if err != nil {
		return fmt.Errorf("Had some kind of error getting the alliance '%s'\n", err)
	}

	response.Alliance = &chremoas_esi.Alliance{
		Id:           request.AllianceId,
		Name:         alliance.AllianceName,
		Ticker:       alliance.Ticker,
		ExecutorCorp: alliance.ExecutorCorp,
		DateFounded:  alliance.DateFounded.Unix(),
	}

	return nil
}

func (eqh *entityQueryHandler) GetAllianceIdCorporations(ctx context.Context, request *chremoas_esi.AllianceRequest, response *chremoas_esi.CorporationsResponse) error {
	alliance, _, err := eqh.ESIClient.ESI.AllianceApi.GetAlliancesAllianceIdCorporations(context.Background(), request.AllianceId, nil)
	if err != nil {
		return fmt.Errorf("Had some kind of error getting the alliance '%s'\n", err)
	}

	response.Corporations = alliance

	return nil
}

func (eqh *entityQueryHandler) GetAllianceIdIcons(ctx context.Context, request *chremoas_esi.AllianceRequest, response *chremoas_esi.AllianceIconResponse) error {
	alliance, _, err := eqh.ESIClient.ESI.AllianceApi.GetAlliancesAllianceIdIcons(context.Background(), request.AllianceId, nil)
	if err != nil {
		return fmt.Errorf("Had some kind of error getting the alliance '%s'\n", err)
	}

	response.Icons = &chremoas_esi.Icons{
		PSixtyFour:   alliance.Px64x64,
		POneTwentyEight: alliance.Px128x128,
	}

	return nil
}

func (eqh *entityQueryHandler) GetAllianceNames(ctx context.Context, request *chremoas_esi.AllianceNamesRequest, response *chremoas_esi.AllianceNamesResponse) error {
	alliance, _, err := eqh.ESIClient.ESI.AllianceApi.GetAlliancesNames(context.Background(), request.AllianceIds, nil)
	if err != nil {
		return fmt.Errorf("Had some kind of error getting the alliance '%s'\n", err)
	}

	for _, v := range alliance {
		response.AllianceNames = append(response.AllianceNames, &chremoas_esi.AllianceNames{AllianceId: v.AllianceId, AllianceName: v.AllianceName})
	}

	return nil
}

func (eqh *entityQueryHandler) GetCorporation(ctx context.Context, request *chremoas_esi.CorporationRequest, response *chremoas_esi.CorporationResponse) error {
	corporation, _, err := eqh.ESIClient.ESI.CorporationApi.GetCorporationsCorporationId(context.Background(), request.EntityId, nil)
	if err != nil {
		return fmt.Errorf("Had some kind of error getting the corporation '%s'\n", err)
	}

	response.Result = &chremoas_esi.Corporation{
		Id:           request.EntityId,
		Name:         corporation.CorporationName,
		Ticker:       corporation.Ticker,
		MemberCount:  corporation.MemberCount,
		CeoId:        corporation.CeoId,
		AllianceId:   corporation.AllianceId,
		Description:  corporation.CorporationDescription,
		TaxRate:      corporation.TaxRate,
		CreationDate: corporation.CreationDate.Unix(),
		CreatorId:    corporation.CreatorId,
		Url:          corporation.Url,
		Faction:      corporation.Faction,
	}

	return nil
}

func (eqh *entityQueryHandler) GetCharacter(ctx context.Context, request *chremoas_esi.CharacterRequest, response *chremoas_esi.CharacterResponse) error {
	character, _, err := eqh.ESIClient.ESI.CharacterApi.GetCharactersCharacterId(context.Background(), request.EntityId, nil)
	if err != nil {
		return fmt.Errorf("Had some kind of error getting the character '%s'\n", err)
	}

	response.Result = &chremoas_esi.Character{
		Id:             request.EntityId,
		Name:           character.Name,
		Description:    character.Description,
		CorporationId:  character.CorporationId,
		AllianceId:     character.AllianceId,
		Birthday:       character.Birthday.Unix(),
		Gender:         character.Gender,
		RaceId:         character.RaceId,
		BloodlineId:    character.BloodlineId,
		AncestryId:     character.AncestryId,
		SecurityStatus: character.SecurityStatus,
		// His docs say this exists, but it doesn't appear to
		//FactionId: character.FactionId
	}

	return nil
}

func (eqh *entityQueryHandler) GetSearch(ctx context.Context, request *chremoas_esi.SearchRequest, response *chremoas_esi.SearchResponse) error {
	// I dislike having to do this as two calls but there is a limit of 10 categories per call
	var categories1 = []string{"agent", "alliance", "character", "constellation", "corporation"}
	result1, _, err1 := eqh.ESIClient.ESI.SearchApi.GetSearch(context.Background(), categories1, request.SearchString, nil)
	if err1 != nil {
		return fmt.Errorf("Had some kind of error searching '%s'\n", err1)
	}

	var categories2 = []string{"faction", "inventorytype", "region", "solarsystem", "station", "wormhole"}
	result2, _, err2 := eqh.ESIClient.ESI.SearchApi.GetSearch(context.Background(), categories2, request.SearchString, nil)
	if err2 != nil {
		return fmt.Errorf("Had some kind of error searching '%s'\n", err2)
	}

	response = &chremoas_esi.SearchResponse{
		Agent:         result1.Agent,
		Alliance:      result1.Alliance,
		Character:     result1.Character,
		Constellation: result1.Constellation,
		Corporation:   result1.Corporation,
		Faction:       result2.Faction,
		Inventorytype: result2.Inventorytype,
		Region:        result2.Region,
		Solarsystem:   result2.Solarsystem,
		Station:       result2.Station,
		Wormhole:      result2.Wormhole,
	}

	return nil
}
