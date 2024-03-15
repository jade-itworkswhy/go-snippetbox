web:
	go run ./cmd/web -addr=":9999"

help:
	go run ./cmd/web -help

download:
	go mod download

verify:
	go mod verify

# will automatically remove any unused packages from your go.mod and go.sum files.
tidy:
	go mod tidy