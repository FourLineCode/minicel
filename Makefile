BIN = minicel
INPUT = input.csv

all:
	go run cmd/main.go $(INPUT)

dev:
	go run cmd/main.go $(filter-out $@,$(MAKECMDGOALS))

build:
	go build -o $(BIN) cmd/main.go

run:
	./$(BIN) $(INPUT)

test: build run

%:
	@: