deps:
	- go mod tidy

build: deps
	- go build -o bin/locksec cmd/main.go