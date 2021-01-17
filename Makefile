all: build start

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

build: deps
	go build -o harvestccode ./main.go

start:
	./harvestccode

lint: deps testdeps
	go vet ./...

clean:
	go clean ./...
