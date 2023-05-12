run:
	go run .
proto:
	protoc --go_out=./api api/api.proto --go-grpc_out=./api

grpcui:
	grpcui -plaintext localhost:8080