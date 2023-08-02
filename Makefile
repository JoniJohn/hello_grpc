.PHONY: run

compile:
	protoc ./proto/*.proto --go_out=./proto/buffs --go-grpc_out=./proto/buffs

test:
	go test -race ./...

run:
	go run main.go