.PHONY: all build install clean

BINARY := oid
INSTALL_DIR := $(HOME)/.local/bin

all: build

build:
	go build -o $(BINARY) .

install: build
	@if [ -d "$(INSTALL_DIR)" ]; then \
		mv -f $(BINARY) "$(INSTALL_DIR)/$(BINARY)"; \
		chmod +x "$(INSTALL_DIR)/$(BINARY)"; \
		echo "installed $(BINARY) to $(INSTALL_DIR)"; \
	else \
		echo "error: $(INSTALL_DIR) does not exist."; \
		exit 1; \
	fi

clean:
	rm -f $(BINARY)