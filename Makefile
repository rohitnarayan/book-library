.PHONY: all mac linux test clean

# Name of the executables
BINARY_NAME_MAC := library_management_system_mac
BINARY_NAME_LINUX := library_management_system_linux

# Source files
SRC := $(wildcard *.go)

all: mac linux test

mac: $(BINARY_NAME_MAC)

linux: $(BINARY_NAME_LINUX)

$(BINARY_NAME_MAC): $(SRC)
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME_MAC) $(SRC)

$(BINARY_NAME_LINUX): $(SRC)
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME_LINUX) $(SRC)

test:
ifeq ($(shell uname),Darwin)
	go test -v ./... -tags mac
else ifeq ($(shell uname),Linux)
	go test -v ./... -tags linux
else
	@echo "Unsupported platform for testing"
	exit 1
endif

clean:
	rm -f $(BINARY_NAME_MAC) $(BINARY_NAME_LINUX)
