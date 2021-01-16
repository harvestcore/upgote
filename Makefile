all: build start

run:
	go run src/main.go

test: testdeps
	go test ./src... -v -count=1 -timeout 30s

deps:
	go get ./src...

testdeps:
	go get -t ./src...

install:
	go install ./src...

build: deps
	go build -o harvestccode ./src/main.go

start:
	./harvestccode

lint: deps testdeps
	go vet ./src...

clean:
	go clean ./src...
