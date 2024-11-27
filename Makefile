# run the api
.PHONY: air
air: swag
	air

.PHONY: api
api: swag
	go run .

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build -trimpath -o ./bin/gin_api_server_template
	# go build -trimpath -o ./bin/gin_api_server_template
	upx ./bin/gin_api_server_template


.PHONY: swag
swag:
	swag init -o ./app/docs
