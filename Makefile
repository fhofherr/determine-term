GO := go

.DEFAULT_GOAL := build

.PHONY: build
build: bin/detterm

bin/detterm: cmd/detterm/main.go
	$(GO) build -o $@ ./$<

.PHONY: install
install:
	cp bin/detterm /usr/local/bin

.PHONY: clean
clean:
	rm -rf bin
