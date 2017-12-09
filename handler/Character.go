package handler

import (
	"fmt"
	"github.com/antihax/goesi"
	"github.com/chremoas/esi-srv/proto"
	"github.com/gregjones/httpcache"
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
)

type characterServiceHandler struct {
	Client    client.Client
	ESIClient *goesi.APIClient
}

func NewCharacterServiceHandler() chremoas_esi.CharacterServiceHandler {
	httpClient := httpcache.NewMemoryCacheTransport().Client()
	return &characterServiceHandler{ESIClient: goesi.NewAPIClient(httpClient, "chremoas-esi-srv Ramdar Chinken on TweetFleet Slack https://github.com/chremoas/esi-srv")}
}

func (eqh *characterServiceHandler) GetCharacterById(ctx context.Context, request *chremoas_esi.GetCharacterByIdRequest, response *chremoas_esi.GetCharacterByIdResponse) error {
	character, _, err := eqh.ESIClient.ESI.CharacterApi.GetCharactersCharacterId(context.Background(), request.Id, nil)
	if err != nil {
		return fmt.Errorf("Had some kind of error getting the character '%s'\n", err)
	}

	response.Character = &chremoas_esi.Character{
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
	}

	return nil
}
