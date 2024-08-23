BINARY=go_audio
VERSION=1.0.0
BUILD_DIR=build
BUILD_TIME=`date +%FT%T%z`
GOX_OS_ARCH="darwin/amd64 darwin/arm64 linux/386 linux/amd64 windows/386 windows/amd64"
GOROOT=/<go_root_path>/bin
CGO_ENABLED=0
CURRENT_PATH=$(shell pwd)

.PHONY: default
default: build

.PHONY: clean
clean:
	rm -rf ${CURRENT_PATH}/${BUILD_DIR}

.PHONY: start
start:
	go run main.go

.PHONY: build
build:
	${GOROOT}/go build -a -o ${CURRENT_PATH}/${BUILD_DIR}/${BINARY} ${CURRENT_PATH}/cmd/main.go

.PHONY: build-version
build-version:
	go build -a -o ${CURRENT_PATH}/${BUILD_DIR}/${BINARY}-${VERSION} ${CURRENT_PATH}/cmd/main.go

.PHONY: build-linux
build-linux:
	GOARCH=amd64 \
	GOOS=linux \
	${GOROOT}/go build -ldflags "-X main.Version=${VERSION}" -a -o ${BUILD_DIR}/${BINARY}-${VERSION} cmd/main.go

.PHONY: build-gox
build-gox:
	gox -ldflags "-X main.Version=${VERSION}" -osarch=${GOX_OS_ARCH} -output="/build/${VERSION}/{{.Dir}}_{{.OS}}_{{.Arch}}"

.PHONY: deps
deps:
	dep ensure;

.PHONY: test-json
test-json:
	${GOROOT}/go test -v ./tests -json > report.json

.PHONY: test
test:
	${GOROOT}/go test -v ./tests > report.txt
