.DEFAULT_GOAL := run

APP = $(shell basename $(CURDIR))
APP_NAME = bf2js.go
BIN_NAME = bf2js

FILE_NAME = $(shell basename $(file))
.PHONY:build
build:
	go build $(APP_NAME)

.PHONY:transpile
transpile:
	./$(BIN_NAME) $(file) > build/$(FILE_NAME).js

.PHONY:run-node
run-node:
	node build/$(FILE_NAME).js

.PHONY:run
run: build transpile run-node



