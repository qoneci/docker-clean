default: build
VERSION ?= 0.0
all:;: '$(VERSION)'

build:
	gox -osarch="linux/amd64 darwin/amd64" -ldflags "-X main.Version=$(VERSION)"

clean:
	@rm docker-clean*

install:
	go install -ldflags "-X main.Version=$(VERSION)"
