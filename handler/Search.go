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

func NewSearchServiceHandler() chremoas_esi.SearchServiceHandler {
	httpClient := httpcache.NewMemoryCacheTransport().Client()
	return &searchServiceHandler{ESIClient: goesi.NewAPIClient(httpClient, "chremoas-esi-srv Ramdar Chinken on TweetFleet Slack https://github.com/chremoas/esi-srv")}
}

func (eqh *searchServiceHandler) Search(ctx context.Context, request *chremoas_esi.SearchRequest, response *chremoas_esi.SearchResponse) error {
	// This doesn't work. Will fix it sometime.
	//var esiOptionals map[string]interface{}
	//
	//foo, _ := ctx.Value("esiOptionals").(map[string]interface{})
	//if len(foo) == 0 {
	//	esiOptionals = nil
	//} else {
	//	esiOptionals = foo
	//}
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

	response.Agent = result1.Agent
	response.Alliance = result1.Alliance
	response.Character = result1.Character
	response.Constellation = result1.Constellation
	response.Corporation = result1.Corporation
	response.Faction = result2.Faction
	response.Inventorytype = result2.InventoryType
	response.Region = result2.Region
	response.Solarsystem = result2.SolarSystem
	response.Station = result2.Station
	//response.Wormhole = result2.Wormhole

	return nil
}
