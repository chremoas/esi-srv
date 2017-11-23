package chremoas_esi

//go:generate protoc --go_out=plugins=micro:. esi-srv.proto
//go:generate mockgen -package=chremoas_esi -source=esi-srv.pb.go -destination=proto_mocks.go EntityQueryClient
