fmt:
	go fmt ./...

lint:
	golangci-lint run

test:
	go test ./... -coverprofile=coverage.out

build:
	go build -o cmd/app ./cmd/app

hooks-install:
	./tools/git-hooks/install.sh
