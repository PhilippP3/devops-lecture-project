FILENAME = main

build:
	go build ./$(SERVICE)/cmd/$(FILENAME).go

run: build
	./$(FILENAME)

clean:
	rm ./$(FILENAME)