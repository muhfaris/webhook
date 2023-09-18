update:
	go mod tidy && go mod vendor -v

dev:
	go run main.go api
