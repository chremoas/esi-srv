package chremoas_esi

//go:generate sh -c "protoc --go_out=plugins=micro:. *.proto"
//go:generate pegomock generate --output=mocks/mock_allianceserviceclient.go --generate-matchers --package=esi_mocks AllianceServiceClient
//go:generate pegomock generate --output=mocks/mock_corporationerviceclient.go --generate-matchers --package=esi_mocks CorporationServiceClient
//go:generate pegomock generate --output=mocks/mock_characterserviceclient.go --generate-matchers --package=esi_mocks CharacterServiceClient
//go:generate pegomock generate --output=mocks/mock_searchserviceclient.go --generate-matchers --package=esi_mocks SearchServiceClient
