APPNAME := "bitbar_bitrise"

.PHONY: build
build:
	go build -o bin/$(APPNAME) .
	chmod +x bin/$(APPNAME)

.PHONY: clean
clean:
	rm -rf bin