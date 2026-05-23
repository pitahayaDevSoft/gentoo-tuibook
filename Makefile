# gentoo-tuibook - Makefile
# Industrial build config compatible with Gentoo Linux and Portage

PREFIX ?= /usr/local
BINDIR = $(PREFIX)/bin
DESTDIR ?=

# Go tools and flags
GO ?= go
GOFLAGS ?=
LDFLAGS ?= -w -s

BINARY = gentoo-tuibook

.PHONY: all build clean install uninstall

all: build

build:
	@echo "Building $(BINARY) nativelly..."
	$(GO) build $(GOFLAGS) -ldflags "$(LDFLAGS)" -o $(BINARY) .

clean:
	@echo "Cleaning up build artifacts..."
	rm -f $(BINARY)

install: build
	@echo "Installing $(BINARY) to $(DESTDIR)$(BINDIR)..."
	install -d $(DESTDIR)$(BINDIR)
	install -m 0755 $(BINARY) $(DESTDIR)$(BINDIR)/$(BINARY)

uninstall:
	@echo "Uninstalling $(BINARY) from $(DESTDIR)$(BINDIR)..."
	rm -f $(DESTDIR)$(BINDIR)/$(BINARY)
