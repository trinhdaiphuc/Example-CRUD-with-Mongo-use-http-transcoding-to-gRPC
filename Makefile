pb:
	protoc -Iprotos/ protos/entity.proto\
		-Ithird_party/googleapis \
		-Ithird_party/grpc-gateway \
		--go_out=plugins=grpc:protos/ \
		--grpc-gateway_out=logtostderr=true:protos/

build:
	- go build -o bin/grpc-service main.go
	- go build -o bin/gateway gateway/main.go

grpc-service:
	- ./bin/grpc-service

gateway-service:
	- ./bin/gateway