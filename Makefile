.PHONY:

test:
	CGO_ENABLED=1 GOARCH=amd64 go test ./... -race --count=1 -cover

mock-prepare:
	go install github.com/golang/mock/mockgen@v1.6.0

mocks:
	mockgen -source=interfaces.go -destination=mock/interfaces.go -package=mock
