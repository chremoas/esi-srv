package esi_mocks

//go:generate sh -c "rm proto_mocks.go"
//go:generate mockgen -package=esi_mocks -destination=proto_mocks.go github.com/chremoas/esi-srv/proto AllianceServiceClient,CharacterServiceClient,CorporationServiceClient,SearchServiceClient
