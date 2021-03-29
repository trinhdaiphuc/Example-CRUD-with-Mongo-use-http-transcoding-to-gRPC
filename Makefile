pb:
	protoc -Iprotos/ protos/*.proto\
		-Ithird_party/googleapis \
		-Ithird_party/protoc-gen-validate \
		--go_out=plugins=grpc:protos/ \
		--grpc-gateway_out=logtostderr=true:protos/ \
		--include_imports --include_source_info \
		--validate_out="lang=go:protos/" \
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