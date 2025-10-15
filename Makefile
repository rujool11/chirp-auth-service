APP_NAME=auth
MAIN_FILE=cmd/auth/main.go
BINARY_PATH=bin/$(APP_NAME)

all: run

build:
	@echo "Building $(APP_NAME)..."
	go build -o $(BINARY_PATH) $(MAIN_FILE)
	@echo "Build complete: $(BINARY_PATH)"

run: build
	@echo "Running $(APP_NAME)..."
	./$(BINARY_PATH)
