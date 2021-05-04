APPNAME := "bitbar_bitrise.2m.cgo"

.PHONY: build
build:
	go build -o bin/$(APPNAME) .
	chmod +x bin/$(APPNAME)

.PHONY: clean
clean:
	rm -rf bin