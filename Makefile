BIN = minicel
FILE = input
INPUT = csv/$(FILE).csv
OS = linux

all:
	go run cmd/main.go $(INPUT) $(ARGS)

build:
	GOOS=$(OS) go build -o $(BIN) cmd/main.go

run:
	./$(BIN) $(INPUT) $(ARGS)

test: build run

compile:
	GOOS=linux go build -o bin/$(BIN)-linux cmd/main.go &&\
	GOOS=windows go build -o bin/$(BIN)-win.exe cmd/main.go