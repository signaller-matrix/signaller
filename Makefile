all: 
	cd cmd; \
	go build -o ../signaller;

install:
	go install github.com/nxshock/signaller/cmd
update-deps:
	go mod tidy
test: update-deps
	go get golang.org/x/tools/cmd/cover
	go get github.com/mattn/goveralls
	go test ./... -v -covermode=count -coverprofile=coverage.out
	$(HOME)/gopath/bin/goveralls -coverprofile=coverage.out -service=travis-ci
