update:
	go mod tidy && go mod vendor -v

dev:
	go run main.go api

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o webhook . && ./webhook
