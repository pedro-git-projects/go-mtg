BINARY := go-mtg 

.PHONY: run build clean

run:
	go run .

build:
	@mkdir -p bin
	go build -o bin/$(BINARY) -v .

clean:
	go clean
	rm -rf bin
