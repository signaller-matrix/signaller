all:
	go test ./...
	go install github.com/nxshock/signaller/cmd
update-deps:
	go mod tidy
