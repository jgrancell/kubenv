TEST?=$$(go list ./... | grep -v vendor)
BINARY=kubenv
VERSION=$$(grep version main.go | head -n1| cut -d\" -f2)
OS_ARCH=linux_amd64
GOBIN=${GOPATH}/bin

default: test

build:
	go build -o ${BINARY}
	chmod +x ${BINARY}

install: build
	mkdir -p ${GOBIN}
	mv ${BINARY} ${GOPATH}/bin/${BINARY}

binaries: build
	rm -rf artifacts
	mkdir -p artifacts
	GOOS=linux GOARCH=amd64 go build -o ${BINARY}-${VERSION}-linux-amd64
	tar -czvf ${BINARY}-${VERSION}-linux-amd64.tar.gz ${BINARY}-${VERSION}-linux-amd64
	mv *.tar.gz artifacts/
	rm -f kubenv*

test:
#	go test -i $(TEST) -cover || exit 1
	echo $(TEST) | xargs -t -n4 go test -timeout=30s -parallel=4 -coverprofile=cover.out
