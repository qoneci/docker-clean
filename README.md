# docker-clean

Why? For fun, and to allow you to have a small process that runs the docker cleanup of un used images/containers/volumes on interval. Written in Go using the native docker api/lib 

tested with golang 1.8

### cli
tested on osx/linux
```bash
usage: docker-clean [<flags>]

Flags:
      --help        Show context-sensitive help (also try --help-long and --help-man).
      --version     Show Version
      --all         Prune all container/images/volumes not used
  -c, --containers  Prune all container not used
  -i, --images      Prune all images not used
  -v, --volumes     Prune all volumes not used
      --demon       run as demon in interval
      --demon-interval=DEMON-INTERVAL
                    demon in interval in SEC, default 300
```

### build deps
```bash
$ brew install glide
$ go get github.com/mitchellh/gox
$ glide up
```

### build with gox
build darwin/amd64 and linux/amd64. use version env to pass version to gox
```bash
VERSION=0.2 make
```
