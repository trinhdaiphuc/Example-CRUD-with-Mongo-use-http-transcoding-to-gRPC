pb:
	protoc -Iprotos/ protos/entity.proto\
		-Ithird_party/googleapis \
		-Ithird_party/grpc-gateway \
		--go_out=plugins=grpc:protos/ \
		--grpc-gateway_out=logtostderr=true:protos/