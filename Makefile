pb:
	protoc -Iprotos/ protos/entity.proto\
		-Ithird_party/googleapis \
		--go_out=plugins=grpc:protos/ \
		--grpc-gateway_out=logtostderr=true:protos/ \
		--include_imports --include_source_info \
		--descriptor_set_out=protos/proto.pb

build:
	- go build -o bin/grpc-service main.go
	- go build -o bin/gateway gateway/main.go

grpc-service:
	- ./bin/grpc-service

gateway-service:
	- ./bin/gateway

dc-gateway:
	- docker-compose up grpc_gateway grpc_app mongo

dc-gateway-build:
	- docker-compose --build up grpc_gateway grpc_app mongo

dc-envoy:
	- docker-compose up envoy-proxy grpc_app mongo 

dc-envoy-build:
	- docker-compose up --build envoy-proxy grpc_app mongo 