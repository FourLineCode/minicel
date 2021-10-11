BIN = minicel
FILE = input
INPUT = csv/$(FILE).csv

all:
	go run cmd/main.go $(INPUT)

build:
	go build -o $(BIN) cmd/main.go

run:
	./$(BIN) $(INPUT)

test: build run