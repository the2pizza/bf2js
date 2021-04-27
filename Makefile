.DEFAULT_GOAL := all

APP = $(shell basename $(CURDIR))
APP_NAME = bf2js.go
BIN_NAME = bf2js

FILE_NAME = $(shell basename $(file))
.PHONY:build
build:
	go build -o bin/$(BIN_NAME) $(APP_NAME)

.PHONY:transpile
transpile:
	./bin/$(BIN_NAME) $(file) > build/$(FILE_NAME).js

.PHONY:run
run:
	node build/$(FILE_NAME).js

.PHONY:all
all: build transpile run



