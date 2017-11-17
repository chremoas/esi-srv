package handler

import (
	"github.com/chremoas/esi-srv/proto"
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
	"github.com/abaeve/auth-srv/proto"
)

type EntityQueryHandler struct {
	Client client.Client
}

func (eqh *EntityQueryHandler) GetAlliances(ctx context.Context, request *chremoas_esi.EntityQueryRequest, response *chremoas_esi.AlliancesResponse) error {
	dbAliances := repository.AllianceRepo.FindAll()
	respAlliances := []*chremoas_esi.Alliance{}

	for _, alliance := range dbAliances {
		respAlliances = append(respAlliances,
			&chremoas_esi.Alliance{
				Id:     alliance.AllianceId,
				Name:   alliance.AllianceName,
				Ticker: alliance.AllianceTicker,
			},
		)
	}

	response.List = respAlliances

	return nil
}

func (eqh *EntityQueryHandler) GetCorporations(ctx context.Context, request *chremoas_esi.EntityQueryRequest, response *chremoas_esi.CorporationsResponse) error {
	dbCorporations := repository.CorporationRepo.FindAll()
	respCorporations := []*chremoas_esi.Corporation{}

	for _, corporation := range dbCorporations {
		respCorporation := &chremoas_esi.Corporation{
			Id:     corporation.CorporationId,
			Name:   corporation.CorporationName,
			Ticker: corporation.CorporationTicker,
		}

		if corporation.AllianceId != nil {
			respCorporation.AllianceId = *corporation.AllianceId
		}

		respCorporations = append(respCorporations, respCorporation)
	}

	response.List = respCorporations

	return nil
}

func (eqh *EntityQueryHandler) GetCharacters(ctx context.Context, request *chremoas_esi.EntityQueryRequest, response *chremoas_esi.CharactersResponse) error {
	dbCharacters := repository.CharacterRepo.FindAll()
	respCharacters := []*chremoas_esi.Character{}

	for _, character := range dbCharacters {
		respCharacters = append(respCharacters,
			&chremoas_esi.Character{
				Id:            character.CharacterId,
				Name:          character.CharacterName,
				CorporationId: character.CorporationId,
			},
		)
	}

	response.List = respCharacters

	return nil
}

func (eqh *EntityQueryHandler) GetEntities(ctx context.Context, request *chremoas_esi.EntityQueryRequest, response *chremoas_esi.EntitiesResponseResponse) error {
	dbRoles := repository.RoleRepo.FindAll()
	respRoles := []*chremoas_esi.Role{}

	for _, role := range dbRoles {
		respRoles = append(respRoles,
			&chremoas_esi.Role{
				RoleName:         role.RoleName,
				ChatServiceGroup: role.ChatServiceGroup,
			},
		)
	}

	response.List = respRoles

	return nil
}
