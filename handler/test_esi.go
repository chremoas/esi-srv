package main

import (
        "fmt"
        "github.com/gregjones/httpcache"
        "github.com/antihax/goesi"
        "golang.org/x/net/context"
)

func main() {
        httpClient := httpcache.NewMemoryCacheTransport().Client()
        client := goesi.NewAPIClient(httpClient, "my esi client http://mysite.com contact <SomeDude> ingame")

        //result, response, err := client.ESI.AllianceApi.GetAlliances(context.Background(), nil)
        //result, _, err := client.ESI.AllianceApi.GetAlliancesAllianceId(context.Background(), 1354830081, nil)

        result, _, err := client.ESI.CorporationApi.GetCorporationsCorporationId(context.Background(), 999, nil)


	if err == nil {
		fmt.Printf("result: %+v\n", result)
	} else {
		fmt.Printf("err: %+v\n", err)
	}
}
