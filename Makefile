ifeq ($(OS),Windows_NT)
BIN=build/bin/aether2.exe
else
BIN=build/bin/aether2
endif
SRC=cmd/aether2
GO_LLVM_DIR=tools/go-llvm

.PHONY: build clean publish

force:

all: run

build:
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

$(BUILDDIR)/aether-linux_amd64.tar: force
	GOOS=linux GOARCH=amd64 go build -o aether-linux_amd64 ./cmd/$(BINNAME)
	tar -cf aether-linux_amd64.tar aether-linux_amd64
	mv aether-linux_amd64.tar $(BUILDDIR)/
	rm aether-linux_amd64

$(BUILDDIR)/aether-linux_386.tar: force
	GOOS=linux GOARCH=386 go build -o aether-linux_386 ./cmd/$(BINNAME)
	tar -cf aether-linux_386.tar aether-linux_386
	mv aether-linux_386.tar $(BUILDDIR)/
	rm aether-linux_386

$(BUILDDIR)/aether-windows_amd64.tar: force
	GOOS=windows GOARCH=amd64 go build -o aether-windows_amd64.exe ./cmd/$(BINNAME)
	tar -cf aether-windows_amd64.tar aether-windows_amd64.exe
	mv aether-windows_amd64.tar $(BUILDDIR)/
	rm aether-windows_amd64.exe

$(BUILDDIR)/aether-windows_386.tar: force
	GOOS=windows GOARCH=386 go build -o aether-windows_386.exe ./cmd/$(BINNAME)
	tar -cf aether-windows_386.tar aether-windows_386.exe
	mv aether-windows_386.tar $(BUILDDIR)/
	rm aether-windows_386.exe
