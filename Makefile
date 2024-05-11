.PHONY: all mac linux clean

# Name of the executables
BINARY_NAME_MAC := library_management_system_mac
BINARY_NAME_LINUX := library_management_system_linux

# Source files
SRC := $(wildcard *.go)

all: mac linux

mac: $(BINARY_NAME_MAC)

linux: $(BINARY_NAME_LINUX)

$(BINARY_NAME_MAC): $(SRC)
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME_MAC) $(SRC)

$(BINARY_NAME_LINUX): $(SRC)
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME_LINUX) $(SRC)

run-mac: $(BINARY_NAME_MAC)
	./$(BINARY_NAME_MAC)

run-linux: $(BINARY_NAME_LINUX)
	./$(BINARY_NAME_LINUX)

clean:
	rm -f $(BINARY_NAME_MAC) $(BINARY_NAME_LINUX)
