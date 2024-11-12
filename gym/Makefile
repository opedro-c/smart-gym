APP_NAME := gym
PKG := ./...
BUILD_DIR := ./bin
ENTRY := ./cmd/main.go

# Load environment variables from .env file
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY: all
all: build run

.PHONY: run
run:
	go run $(ENTRY)

.PHONY: build
build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(ENTRY)

.PHONY: debug
debug:
	if [ ! -d "$(BUILD_DIR)" ]; then \
		make build; \
	fi
	gdb $(BUILD_DIR)/$(APP_NAME)

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)/*

.PHONY: migrate
migrate:
	for file in internal/database/sqlite/migrations/*.sql; do \
		sqlite3 gym.db < $$file; \
	done

.PHONY: help
help:
	@echo "Makefile for Gym service"
	@echo ""
	@echo "Usage:"
	@echo "  make run        Run the application"
	@echo "  make build      Build the application"
	@echo "  make debug      Debug the application using GDB"
	@echo "  make clean      Clean the build directory"
