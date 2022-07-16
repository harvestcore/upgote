all: buildapp start

run:
	go run ./main.go

test: testdeps
	go test ./... -v -count=1 -timeout 30s

deps:
	go get ./...

testdeps:
	go get -t ./...

install:
	go install ./...

build:
	echo "Building app..."

buildapp: deps
	go build -o upgote ./main.go

start:
	./upgote

lint: deps testdeps
	go vet ./...

clean:
	go clean ./...
