FILENAME = main

build:
	go build ./cmd/$(FILENAME).go

run: build
	./$(FILENAME)

clean:
	rm ./$(FILENAME)