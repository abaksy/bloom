# Variables
BINARY_NAME=bloomfilter
BINARY_DIR=bin
CMD_DIR=cmd

# Ensure bin directory exists
$(BINARY_DIR):
	mkdir -p $(BINARY_DIR)

# Build the project
build: $(BINARY_DIR)
	go build -o $(BINARY_DIR)/$(BINARY_NAME) ./$(CMD_DIR)/main.go

# Clean build artifacts
clean:
	rm -rf $(BINARY_DIR)

# Run the project
run:
	go run ./$(CMD_DIR)/main.go

# Build and run
all: build run

.PHONY: build clean run all