all:
	go test ./...
	cd cmd
	go build
update-deps:
	go mod tidy
