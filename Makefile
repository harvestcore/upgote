run:
	go run src/main.go

test:
	go test ./tests...

install:
	cd src && go get ./...