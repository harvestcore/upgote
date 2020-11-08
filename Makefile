run:
	cd src && go run main.go

test:
	cd src && go test ./...

install:
	cd src && go get ./...

lint:
	cd src && go vet