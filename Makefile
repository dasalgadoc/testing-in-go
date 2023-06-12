test:
	go test -v ./...

run:
	go run main.go

all: test run
