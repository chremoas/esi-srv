package chremoas_esi

//go:generate protoc --go_out=plugins=micro:. alliance.proto
//go:generate protoc --go_out=plugins=micro:. character.proto
//go:generate protoc --go_out=plugins=micro:. corporation.proto
//go:generate protoc --go_out=plugins=micro:. search.proto
