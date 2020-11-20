all: install

run:
	go run src/main.go

test: testdeps
	go test ./src... -v

deps:
	go get ./src...

testdeps:
	go get -t ./src...

install:
	go install ./src...

build:
	go build ./src...

lint:
	go vet ./src...

clean:
	go clean ./src...
