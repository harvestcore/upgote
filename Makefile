all: install

run:
	go run src/main.go

test: testdeps
	go test ./src... -v -count=1

deps:
	go get ./src...

testdeps:
	go get -t ./src...

install:
	go install ./src...

build:
	go build ./src...

lint: deps testdeps
	go vet ./src...

clean:
	go clean ./src...
