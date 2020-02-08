# Example project using gRPC and http transcoding

## Install tools

Install dep tool: https://golang.github.io/dep/docs/installation.html  
Install MongoDB and make sure it's running on localhost:27017

## Make sure all the dependencies is in sync

`dep status` && `dep ensure`

## Generate gRPC stub
- Generating client and server code

  protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=Mgoogle/api/annotations.proto=github.com/gengo/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. \
  protos/entity.proto

- Generate reverse-proxy for your RESTful API:

  protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:. \
  protos/entity.proto

## Start Server

`go run server/*.go`
`go run gateway/main.go`

## Example API Calls

## List entities

`curl -X GET 'http://localhost:8080/entities'`

## Create entity

`curl -X POST 'http://localhost:8080/entities' -d '{"name":"Phuc qua dep trai","description":"Kha la banh","url":"phucdeptrai.com.vn"}'`

## Read entity

`curl -X GET "http://localhost:8080/entities/5d11e96b9dadaf6eef8599be"`

## Update entity

`curl -X PUT 'http://localhost:8080/entities' -d '{"id":"5dff0ab0ac327677d38754dd","name":"Phuc dep trai vai","description":"Qua la banh","url":"phuchotboy.com.vn"}'`

## Delete entity

`curl -X DELETE "http://localhost:8080/entities/5d11e8ee9dadaf6eef8599b9"`

## USE KONG AS API GATEWAY

- If you want to use Kong (https://konghq.com/kong/) as API gateway. You can checkout to banch kong-api-gw

`git checkout kong-api-gw`
