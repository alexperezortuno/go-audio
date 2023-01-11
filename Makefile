BINARY=go_auth
VERSION=1.0.0
BUILD_DIR=./build
BUILD_TIME=`date +%FT%T%z`
GOX_OS_ARCH="darwin/amd64 darwin/arm64 linux/386 linux/amd64 windows/386 windows/amd64"

.PHONY: default
default: build

.PHONY: clean
clean:
	rm -rf ./build

.PHONY: start
start:
	CGO_ENABLED=0 \
	go run main.go

.PHONY: build
build:
	CGO_ENABLED=0 \
	go build -a -o ${BUILD_DIR}/${BINARY} main.go

.PHONY: build-version
build-version:
	CGO_ENABLED=0 \
	go build -a -o ${BUILD_DIR}/${BINARY}-${VERSION} main.go

.PHONY: build-linux
build-linux:
	CGO_ENABLED=0 \
	GOARCH=amd64 \
	GOOS=linux \
	go build -ldflags "-X main.Version=${VERSION}" -a -o ${BUILD_DIR}/${BINARY}-${VERSION} main.go

.PHONY: build-gox
build-gox:
	gox -ldflags "-X main.Version=${VERSION}" -osarch=${GOX_OS_ARCH} -output="/build/${VERSION}/{{.Dir}}_{{.OS}}_{{.Arch}}"

.PHONY: deps
deps:
	dep ensure;

.PHONY: test
test:
	go test
