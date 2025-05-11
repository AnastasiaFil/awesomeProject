build:
	go mod download && go build ./cmd/main.go

run: build
	docker-compose up --remove-orphans

test:
	go test -v ./...

swag:
	swag init -g cmd/main.go
