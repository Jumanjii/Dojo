SOURCES = $(shell find . -path ./vendor -prune -o -name "*.go" -print)

all: bin/server

bin/server: $(SOURCES)
	GO_ENABLED=0 GOOS=linux go build -tags netgo -o bin/server main.go

run:
	bin/server

local: local/server

local/server: $(SOURCES)
	GO_ENABLED=0 go build -tags netgo -o bin/server main.go

.PHONY: all local run
