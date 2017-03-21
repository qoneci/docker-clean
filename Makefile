default: build
VERSION ?= 0.0
all:;: '$(VERSION)'

build:
	gox -osarch="linux/amd64 darwin/amd64" -ldflags "-X main.Version=$(VERSION)"

clean:
	@rm docker-clean_* ubuntu-xenial-16.04-cloudimg-console.log
