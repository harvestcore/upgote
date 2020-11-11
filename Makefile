all: install build

run:
	cd src && go run main.go

test:
	cd src && go test ./...

install:
	cd src && go get ./...

build:
	cd src && go build

lint:
	cd src && go vet
