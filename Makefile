BINARY_NAME=godfs
OS_NAME=$(shell uname -s | tr A-Z a-z)
OS_ARCH=$(shell uname -m | tr A-Z a-z)

build:
	GOARCH=${OS_ARCH} GOOS=darwin go build -o bin/${BINARY_NAME}-${OS_NAME}

run: build
	./bin/${BINARY_NAME}-${OS_NAME}

clean:
	go clean
	rm bin/${BINARY_NAME}-darwin
	rm bin/${BINARY_NAME}-linux
	rm bin/${BINARY_NAME}-windows