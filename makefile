web:
	go run ./cmd/web -addr=":4000"

help:
	go run ./cmd/web -help

download:
	go mod download

verify:
	go mod verify

# will automatically remove any unused packages from your go.mod and go.sum files.
tidy:
	go mod tidy 

test-all:
	go test ./...

test-no-cache:
	go test -count=1 ./...

test-web:
	go test ./cmd/web -v