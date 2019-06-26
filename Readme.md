##-----GENERATE FROM PROTOBUFFER FILE-----##
## Generating client and server code
protoc -I/usr/local/include -I.  -I$GOPATH/src  -Ithird_party/googleapis  --go_out=plugins=grpc:. server/entity/entity.proto 
## Generate reverse-proxy for your RESTful API
protoc -I/usr/local/include -I. -I$GOPATH/src -Ithird_party/googleapis --grpc-gateway_out=logtostderr=true:. server/entity/entity.proto

##-----GET LIBRARIES----------------------##
go get

##-----RUN SERVER-------------------------##
## Run server gRPC
go run server/*.go
## Run server http 
go run proxy/main-rproxy.go

##-------Some example tests----------------##
## List entities
curl -X GET 'http://localhost:8080/entities'
## Create entity
curl -X POST 'http://localhost:8080/entities' \
-d '{"name":"Phuc qua dep trai","description":"Kha la banh","url":"phucdeptrai.com.vn"}'
## Read entity
curl -X GET "http://localhost:8080/entities/5d11e96b9dadaf6eef8599be"
## Update entity
curl -X PUT 'http://localhost:8080/entities' \
-d '{"id":"5d11e8ee9dadaf6eef8599b9","name":"Phuc qua dep trai","description":"Kha la banh","url":"phucdeptrai.com.vn"}'
## Delete entity
curl -X DELETE "http://localhost:8080/entities/5d11e8ee9dadaf6eef8599b9"