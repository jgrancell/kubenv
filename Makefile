TEST?=$$(go list ./... | grep -v vendor)
BINARY=kubenv
VERSION=0.1.0
OS_ARCH=linux_amd64
GOBIN=$GOPATH/bin

default: test

build:
	go build -o ${BINARY}
	chmod +x ${BINARY}

install: build
	mkdir -p $GOBIN
	mv ${BINARY} ${GOPATH}/bin/${BINARY}

test:
#	go test -i $(TEST) -cover || exit 1
	echo $(TEST) | xargs -t -n4 go test -timeout=30s -parallel=4 -coverprofile=cover.out
