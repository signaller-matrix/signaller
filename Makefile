all:
	go test ./...
	go build github.com/nxshock/signaller/cmd
update-deps:
	go mod tidy
