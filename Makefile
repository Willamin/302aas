.PHONY: build
build:
	mkdir -p functions
	go get ./...
	go build -o functions/302aas ./...

.PHONY: run
run:
	./functions/302aas
