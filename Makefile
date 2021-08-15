.DEFAULT_GOAL := dc-all

build:
	go build -o bin/gateway gateway/main.go
	go build -o bin/denny denny/main.go
	go build -o bin/grpc-server main.go
	go build -o bin/gateway-server gateway-server/main.go

install-proto-gen-go:
	go get -u github.com/golang/protobuf/protoc-gen-go

install-third-party:
	git submodule update --init

install-protoc-gen-swagger:
	go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2

install-protoc-gen-grpc-gateway:
	go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway

install: install-proto-gen-go install-third-party install-protoc-gen-swagger install-protoc-gen-grpc-gateway

gen-protobuf:
	protoc -Iprotos/ protos/*.proto \
		-Ithird_party/googleapis \
		-Ithird_party/protoc-gen-validate \
		--go_out=plugins=grpc:protos/ \
		--grpc-gateway_out=logtostderr=true:protos/ \
		--validate_out="lang=go:protos/" \
		--swagger_out=logtostderr=true:public \
		--include_imports --include_source_info \
		--descriptor_set_out=protos/proto.pb

dc-gateway:
	- docker-compose up grpc_gateway grpc_app mongo

dc-gateway-build:
	- docker-compose --build up grpc_gateway grpc_app mongo

dc-envoy:
	- docker-compose up envoy-proxy grpc_app mongo 

dc-envoy-build:
	- docker-compose up --build envoy-proxy grpc_app mongo

dc-denny:
	- docker-compose up denny mongo

dc-denny-build:
	- docker-compose up --build denny mongo

dc-all: gen-protobuf
	- docker-compose up