
.PHONY: run

SOURCES = $(shell find . -name '*.go')

run: $(SOURCES)
	@go run cmd/main.go
