export GOOS=windows

fmt:
	gofmt -l -w -s .

lint:
	gometalinter --vendor --config gometalinter.config ./...

test:
	go test -v ./...

build:
	promu build -v
