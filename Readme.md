##-----GET LIBRARIES----------------------##
## Run this command line
go get

##-----RUN SERVER-------------------------##
## Run server gRPC
go run server/*.go
## Run server http 
go run gateway/main.go

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

##-------USE KONG AS API GATEWAY--------------##
## If you want to use Kong (https://konghq.com/kong/) as API gateway. You can checkout to banch kong-api-gw
git checkout kong-api-gw