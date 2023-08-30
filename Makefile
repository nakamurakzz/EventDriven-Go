APP_NAME=go-iot-gateway

air:
	@air

build:
	@go build -o bin/$(APP_NAME) main.go

start:build
	@./bin/$(APP_NAME)
