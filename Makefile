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

generate-certs:
	openssl req -x509 -nodes -newkey rsa:2048 -days 365 -keyout privkey.pem -out cert.pem -subj "/CN=127.0.0.1"
	openssl  x509 -in cert.pem -out root.crt