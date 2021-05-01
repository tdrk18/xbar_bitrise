APPNAME := "bitbar_bitrise"

.PHONY: build
build:
	go build -o bin/$(APPNAME) .

.PHONY: clean
clean:
	rm -rf bin