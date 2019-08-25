GO := go

.DEFAULT_GOAL := build

.PHONY: build
build: bin/determine-term

bin/determine-term: cmd/determine-term/main.go
	$(GO) build -o $@ ./$<

.PHONY: install
install:
	cp bin/determine-term /usr/local/bin

.PHONY: clean
clean:
	rm -rf bin
