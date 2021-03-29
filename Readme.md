# Example project using gRPC and http transcoding

> This project is an example using [gRPC gateway](https://grpc-ecosystem.github.io/grpc-gateway/), 
> [Envoy proxy](https://www.envoyproxy.io/) to transcode gRPC API server into REST API server. Furthermore, I use 
> [denny framework](https://github.com/whatvn/denny) which is not transcode gRPC API server into REST API server. Denny 
> not only exposes gRPC and http api server in 1 port but also invoke the code you wrote in grpc functions, does not 
> trigger grpc call when you call http.

## Generate gRPC stub

Generating client and server code and reverse-proxy for your REST API:

```shell
make gen-protobuf
```

## Start project

- Run all project

```shell
make dc-all
```

- Use grpc gateway

```shell
make dc-gateway
```

- Use envoy

```shell
make dc-envoy
```

- Use denny

```shell
make dc-denny
```

## Example API Calls:

> Install extension [REST client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) on VS Code for 
> easy send api call with ***.http** files in **/client** folder

- List entities

```shell
curl -X GET 'http://localhost:8080/entities'
```

- Create entity

```shell
curl -X POST 'http://localhost:8080/entities' -d '{"name":"Phuc qua dep trai","description":"Kha la banh","url":"phucdeptrai.com.vn"}'
```

- Read entity

```shell
curl -X GET "http://localhost:8080/entities/5d11e96b9dadaf6eef8599be"
```

- Update entity

```shell
curl -X PUT 'http://localhost:8080/entities' -d '{"id":"5dff0ab0ac327677d38754dd","name":"Phuc dep trai vai","description":"Qua la banh","url":"phuchotboy.com.vn"}'
```

- Delete entity

```shell
curl -X DELETE "http://localhost:8080/entities/5d11e8ee9dadaf6eef8599b9"
```
