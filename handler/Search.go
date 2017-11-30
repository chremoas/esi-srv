package handler

import (
	"fmt"
	"github.com/antihax/goesi"
	"github.com/chremoas/esi-srv/proto"
	"github.com/gregjones/httpcache"
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
)

type searchServiceHandler struct {
	Client    client.Client
	ESIClient *goesi.APIClient
}

func NewSearchQueryHandler() chremoas_esi.SearchServiceHandler {
	httpClient := httpcache.NewMemoryCacheTransport().Client()
	return &searchServiceHandler{ESIClient: goesi.NewAPIClient(httpClient, "chremoas-esi-srv Ramdar Chinken on TweetFleet Slack https://github.com/chremoas/esi-srv")}
}

func (eqh *searchServiceHandler) Search(ctx context.Context, request *chremoas_esi.SearchRequest, response *chremoas_esi.SearchResponse) error {
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
