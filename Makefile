ifeq ($(OS),Windows_NT)
BIN=bin/aether2.exe
else
BIN=bin/aether2
endif
SRC=bin
GO_LLVM_DIR=tools/go-llvm

.PHONY: build clean deps

deps:
	@echo "Checking for mold..."
	@which mold || which mold.exe || (echo "mold not found. Installing..." && \
		(which apt && sudo apt update && sudo apt install -y mold) || \
		(which brew && brew install mold) || \
		(which pacman && sudo pacman -S mold) || \
		(which dnf && sudo dnf install mold) || \
		(which yum && sudo yum install mold) || \
		(which zypper && sudo zypper install mold) || \
		(echo "No supported package manager found. Please install mold manually." && exit 1))
	@echo "Checking for LLVM development libraries..."
	@pkg-config --exists llvm || (echo "LLVM dev libraries not found. Installing..." && \
		(which apt && sudo apt install -y llvm-dev) || \
		(which brew && brew install llvm) || \
		(which pacman && sudo pacman -S llvm) || \
		(which dnf && sudo dnf install llvm-devel) || \
		(which yum && sudo yum install llvm-devel) || \
		(which zypper && sudo zypper install llvm-devel) || \
		(echo "No supported package manager found. Please install LLVM dev libraries manually." && exit 1))

all: run

build: deps
	CGO_ENABLED=1 go build -o $(BIN) ./$(SRC)

# Usage: make run ARGS="test.ae"
run:
	go run $(SRC) $(ARGS)

test:
	go test -v ./...

clean:
	go clean -cache -testcache


PLATFORMS = linux windows
ARCHS = amd64 386
BINNAME = aether2
BUILDDIR = build/bin

publish: $(BUILDDIR) $(foreach p,$(PLATFORMS),$(foreach a,$(ARCHS),$(BUILDDIR)/aether-$(p)_$(a).tar))

$(BUILDDIR):
	mkdir -p $(BUILDDIR)

$(BUILDDIR)/aether-linux_amd64.tar:
	GOOS=linux GOARCH=amd64 go build -o aether-linux_amd64 ./cmd/$(BINNAME)
	tar -cf aether-linux_amd64.tar aether-linux_amd64
	mv aether-linux_amd64.tar $(BUILDDIR)/
	rm aether-linux_amd64

$(BUILDDIR)/aether-linux_386.tar:
	GOOS=linux GOARCH=386 go build -o aether-linux_386 ./cmd/$(BINNAME)
	tar -cf aether-linux_386.tar aether-linux_386
	mv aether-linux_386.tar $(BUILDDIR)/
	rm aether-linux_386

$(BUILDDIR)/aether-windows_amd64.tar:
	GOOS=windows GOARCH=amd64 go build -o aether-windows_amd64.exe ./cmd/$(BINNAME)
	tar -cf aether-windows_amd64.tar aether-windows_amd64.exe
	mv aether-windows_amd64.tar $(BUILDDIR)/
	rm aether-windows_amd64.exe

$(BUILDDIR)/aether-windows_386.tar:
	GOOS=windows GOARCH=386 go build -o aether-windows_386.exe ./cmd/$(BINNAME)
	tar -cf aether-windows_386.tar aether-windows_386.exe
	mv aether-windows_386.tar $(BUILDDIR)/
	rm aether-windows_386.exe