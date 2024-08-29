run:
	go run cmd/cacheGo/main.go

test:
	gotest -v ./...

coverage:
	go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

