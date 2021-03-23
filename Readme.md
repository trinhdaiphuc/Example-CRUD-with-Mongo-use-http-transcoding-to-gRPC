# Example project using gRPC and http transcoding

> This project is an example using gRPC gateway and Envoy proxy to transcode gRPC API server into REST API.

## Generate gRPC stub

Generating client and server code and reverse-proxy for your RESTful API:

```shell
make pb
```

## Start project

- Use grpc gateway

```shell
make dc-gateway
```

- Use envoy

```shell
make dc-envoy
```

## Example API Calls:

_Install extension [REST client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) on VS Code for easy send api call in **rest.http** file_

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
