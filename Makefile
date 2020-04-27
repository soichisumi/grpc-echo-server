protoc:
	protoc --go_out=plugins=grpc:./pkg ./proto/*.proto

go-build:
	go build -o exe -mod vendor .

vendor:
	go mod vendor

go-run:
	go run .

run-msgencoder:
	go run ./scripts/msg-encoder/