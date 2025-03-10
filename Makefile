FILENAME = main

build:
	go build .\cmd\$(FILENAME).go

run: build
	./$(FILENAME).exe