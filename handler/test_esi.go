package main

import (
        "fmt"
        "github.com/gregjones/httpcache"
        "github.com/antihax/goesi"
        "golang.org/x/net/context"
	"strings"
	"os"
)

func main() {
        httpClient := httpcache.NewMemoryCacheTransport().Client()
        client := goesi.NewAPIClient(httpClient, "my esi client http://mysite.com contact <SomeDude> ingame")

        //result, response, err := client.ESI.AllianceApi.GetAlliances(context.Background(), nil)
        //result, _, err := client.ESI.AllianceApi.GetAlliancesAllianceId(context.Background(), 1354830081, nil)

        //result, _, err := client.ESI.CorporationApi.GetCorporationsCorporationId(context.Background(), 999, nil)
	search := strings.Join(os.Args[1:], " ")
	var things = []string{"alliance", "corporation", "character"}
        result, _, err := client.ESI.SearchApi.GetSearch(context.Background(), things, search, nil)

	if len(result.Corporation) != 0 {
		fmt.Printf("Corp: %+v\n", result.Corporation)
	} else if len(result.Alliance) != 0 {
		fmt.Printf("Alliance: %+v\n", result.Alliance)
	} else if len(result.Character) != 0 {
		fmt.Printf("Character: %+v\n", result.Character)
	} else {
		fmt.Println("WTF!\n")
	}

	if err != nil {
		fmt.Printf("err: %+v\n", err)
	} else {
		fmt.Printf("result: %+v\n", result)
	}
}
