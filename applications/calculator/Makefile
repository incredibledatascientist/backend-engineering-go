APP=calculator
BIN_DIR=bin
BIN=$(BIN_DIR)/$(APP)

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN)

run: build
	$(BIN)

clean:
	rm -rf $(BIN_DIR) logs

fmt:
	go fmt ./...

test:
	go test ./...
