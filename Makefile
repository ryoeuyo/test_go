APPNAME := test_go

build:
	go build -o ./bin/$(APPNAME) ./cmd/main.go

run: build
	./bin/$(APPNAME)

test:
	go test -v ./internal/http/rest/
