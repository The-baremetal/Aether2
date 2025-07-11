# 🍕 Aether CMake Makefile
# Simple shortcuts for common operations

.PHONY: all build debug release clean install test format lint docs help

# Default target
all: build

# Build targets
build:
	@echo "🍕 Building Aether CMake Demo..."
	@mkdir -p build
	@cd build && cmake .. && make

debug:
	@echo "🍕 Building Debug version..."
	@mkdir -p build-debug
	@cd build-debug && cmake -DCMAKE_BUILD_TYPE=Debug .. && make

release:
	@echo "🍕 Building Release version..."
	@mkdir -p build-release
	@cd build-release && cmake -DCMAKE_BUILD_TYPE=Release .. && make

# Development targets
format:
	@echo "🍕 Formatting Aether code..."
	@aether format src/

lint:
	@echo "🍕 Linting Aether code..."
	@aether lint src/

test:
	@echo "🍕 Running Aether tests..."
	@aether test

docs:
	@echo "🍕 Generating documentation..."
	@aether docs

# Installation
install: build
	@echo "🍕 Installing Aether Demo..."
	@cd build && make install

# Package creation
package: build
	@echo "🍕 Creating package..."
	@cd build && cpack

# Cross-compilation
cross-linux:
	@echo "🍕 Cross-compiling for Linux..."
	@mkdir -p build-linux
	@cd build-linux && cmake -DCMAKE_AETHER_TARGET_OS=linux -DCMAKE_AETHER_TARGET_ARCH=amd64 .. && make

cross-windows:
	@echo "🍕 Cross-compiling for Windows..."
	@mkdir -p build-windows
	@cd build-windows && cmake -DCMAKE_AETHER_TARGET_OS=windows -DCMAKE_AETHER_TARGET_ARCH=amd64 .. && make

cross-macos:
	@echo "🍕 Cross-compiling for macOS..."
	@mkdir -p build-macos
	@cd build-macos && cmake -DCMAKE_AETHER_TARGET_OS=darwin -DCMAKE_AETHER_TARGET_ARCH=amd64 .. && make

# Library operations
lib-shared:
	@echo "🍕 Building shared library..."
	@aether build --create-library --library-type shared --library-name mathlib src/mathlib.ae

lib-static:
	@echo "🍕 Building static library..."
	@aether build --create-library --library-type static --library-name mathlib src/mathlib.ae

lib-both:
	@echo "🍕 Building both library types..."
	@aether build --create-library --library-type both --library-name mathlib src/mathlib.ae

# Analysis
analyze:
	@echo "🍕 Analyzing Aether code..."
	@aether build --analyze-only src/

deps:
	@echo "🍕 Checking dependencies..."
	@aether deps

# Cleaning
clean:
	@echo "🍕 Cleaning build directories..."
	@rm -rf build build-debug build-release build-linux build-windows build-macos
	@rm -f *.so *.a *.dll *.lib *.pc

clean-all: clean
	@echo "🍕 Deep cleaning..."
	@find . -name "*.o" -delete
	@find . -name "*.ll" -delete
	@find . -name "*.bc" -delete
	@find . -name "*.s" -delete

# Help
help:
	@echo "🍕 Aether CMake Demo - Available targets:"
	@echo ""
	@echo "Build targets:"
	@echo "  build     - Build with default settings"
	@echo "  debug     - Build with debug information"
	@echo "  release   - Build with release optimizations"
	@echo ""
	@echo "Development:"
	@echo "  format    - Format Aether source code"
	@echo "  lint      - Lint Aether source code"
	@echo "  test      - Run Aether tests"
	@echo "  docs      - Generate documentation"
	@echo ""
	@echo "Library operations:"
	@echo "  lib-shared - Build shared library"
	@echo "  lib-static - Build static library"
	@echo "  lib-both   - Build both library types"
	@echo ""
	@echo "Cross-compilation:"
	@echo "  cross-linux  - Cross-compile for Linux"
	@echo "  cross-windows - Cross-compile for Windows"
	@echo "  cross-macos  - Cross-compile for macOS"
	@echo ""
	@echo "Analysis:"
	@echo "  analyze - Analyze Aether code"
	@echo "  deps    - Check dependencies"
	@echo ""
	@echo "Installation:"
	@echo "  install  - Install the application"
	@echo "  package  - Create distribution package"
	@echo ""
	@echo "Cleaning:"
	@echo "  clean     - Clean build directories"
	@echo "  clean-all - Deep clean everything"
	@echo ""
	@echo "Help:"
	@echo "  help - Show this help message"

# Default target
.DEFAULT_GOAL := help 