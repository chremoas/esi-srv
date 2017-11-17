package handler

import (
	"github.com/antihax/goesi"
	"github.com/chremoas/esi-srv/proto"
	"github.com/gregjones/httpcache"
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
	"fmt"
)

type entityQueryHandler struct {
	Client    client.Client
	ESIClient *goesi.APIClient
}

func NewEntityQueryHandler() chremoas_esi.EntityQueryHandler {
	httpClient := httpcache.NewMemoryCacheTransport().Client()
	return &entityQueryHandler{ESIClient: goesi.NewAPIClient(httpClient, "chremoas-esi-srv Ramdar Chinken on TweetFleet Slack https://github.com/chremoas/esi-srv")}
}

func (eqh *entityQueryHandler) GetAlliance(ctx context.Context, request *chremoas_esi.EntityQueryRequest, response *chremoas_esi.AlliancesResponse) error {
	alliance, _, err := eqh.ESIClient.ESI.AllianceApi.GetAlliancesAllianceId(context.Background(), corporation.AllianceId, nil)
	if err != nil {
		return fmt.Errorf("Had some kind of error getting the alliance '%s'\n", err)
	}

	//dbAliances := repository.AllianceRepo.FindAll()
	//respAlliances := []*chremoas_esi.Alliance{}
	//
	//for _, alliance := range dbAliances {
	//	respAlliances = append(respAlliances,
	//		&chremoas_esi.Alliance{
	//			Id:     alliance.AllianceId,
	//			Name:   alliance.AllianceName,
	//			Ticker: alliance.AllianceTicker,
	//		},
	//	)
	//}
	//
	//response.List = respAlliances

	return nil
}

func (eqh *entityQueryHandler) GetCorporation(ctx context.Context, request *chremoas_esi.EntityQueryRequest, response *chremoas_esi.CorporationsResponse) error {
	corporation, _, err := eqh.ESIClient.ESI.CorporationApi.GetCorporationsCorporationId(context.Background(), character.CorporationId, nil)
	if err != nil {
		return fmt.Errorf("Had some kind of error getting the corporation '%s'\n", err)
	}

	//dbCorporations := repository.CorporationRepo.FindAll()
	//respCorporations := []*chremoas_esi.Corporation{}
	//
	//for _, corporation := range dbCorporations {
	//	respCorporation := &chremoas_esi.Corporation{
	//		Id:     corporation.CorporationId,
	//		Name:   corporation.CorporationName,
	//		Ticker: corporation.CorporationTicker,
	//	}
	//
	//	if corporation.AllianceId != nil {
	//		respCorporation.AllianceId = *corporation.AllianceId
	//	}
	//
	//	respCorporations = append(respCorporations, respCorporation)
	//}
	//
	//response.List = respCorporations

	return nil
}

func (eqh *entityQueryHandler) GetCharacter(ctx context.Context, request *chremoas_esi.EntityQueryRequest, response *chremoas_esi.CharactersResponse) error {
	character, _, err := eqh.ESIClient.ESI.CharacterApi.GetCharactersCharacterId(context.Background(), int32(verifyReponse.CharacterID), nil)
	if err != nil {
		return fmt.Errorf("Had some kind of error getting the character '%s'\n", err)
	}

	//dbCharacters := repository.CharacterRepo.FindAll()
	//respCharacters := []*chremoas_esi.Character{}
	//
	//for _, character := range dbCharacters {
	//	respCharacters = append(respCharacters,
	//		&chremoas_esi.Character{
	//			Id:            character.CharacterId,
	//			Name:          character.CharacterName,
	//			CorporationId: character.CorporationId,
	//		},
	//	)
	//}
	//
	//response.List = respCharacters

	return nil
}

func (eqh *entityQueryHandler) GetEntity(ctx context.Context, request *chremoas_esi.EntityQueryRequest, response *chremoas_esi.EntitiesResponseResponse) error {

	//dbRoles := repository.RoleRepo.FindAll()
	//respRoles := []*chremoas_esi.Role{}
	//
	//for _, role := range dbRoles {
	//	respRoles = append(respRoles,
	//		&chremoas_esi.Role{
	//			RoleName:         role.RoleName,
	//			ChatServiceGroup: role.ChatServiceGroup,
	//		},
	//	)
	//}
	//
	//response.List = respRoles

	return nil
}
