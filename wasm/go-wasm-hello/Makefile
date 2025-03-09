# Variables
BINARY_NAME=main.wasm
SRC=./...

# Default target
all: build

# Build the Go binary
build:
	GOOS=wasip1 GOARCH=wasm go build -o $(BINARY_NAME) $(SRC)

# Run the Go application
run: build
	wasmtime run ./$(BINARY_NAME)

# Clean up generated files
clean:
	rm -f $(BINARY_NAME)

# Run tests
 test:
	go test -v $(SRC)

# Format the code
fmt:
	go fmt $(SRC)

.PHONY: all build run clean test fmt
