# Aether Installation Guide

## Introduction

Aether is a modern programming language and toolchain. This guide describes how to install all prerequisites and build Aether from source on Linux, macOS, and Windows.

## Prerequisites

You must have the following tools installed:

- Go (>=1.20)
- LLVM (>=15)
- aria2
- mold linker

### Checking for Prerequisites

#### Go

- Check: `go version`
- Install:
  - Linux: `sudo apt install golang` or `sudo dnf install golang` or `sudo pacman -Sy go`
  - macOS: `brew install go`
  - Windows: [Download from golang.org](https://golang.org/dl/)

#### LLVM

- Check: `llvm-config --version`
- Install: Run `scripts/prerequisites/llvm.sh` (Linux or WSL only)
  - macOS: `brew install llvm`
  - Windows: [Download from LLVM releases](https://releases.llvm.org/)

#### aria2

- Check: `aria2c --version`
- Install: Run `scripts/prerequisites/aria2.sh`

#### mold linker

- Check: `mold --version`
- Install: Run `scripts/prerequisites/mold.sh`

## Running Prerequisite Scripts

On Linux/macOS, run the following scripts as needed:

```bash
bash scripts/prerequisites/llvm.sh
bash scripts/prerequisites/aria2.sh
bash scripts/prerequisites/mold.sh
```

On Windows, use Chocolatey or Scoop, or install manually as described above.

## Building Aether

1. Clone the repository:

   ```bash
   git clone https://github.com/The-baremetal/Aether2.git
   cd Aether2
   ```

2. Build using Make:

   ```bash
   make
   ```

   ```bash
   Or build with Go:
   ```

   ```bash
   go build ./...
   ```

## Testing

To verify your build, run:

```bash
go test ./...
```

Or use the Aether CLI:

```bash
./aether2 --version
```

## Updating Aether

To update Aether to the latest version, follow these steps:

1. Pull the latest changes from the repository:

   ```bash
   git pull origin main
   ```

2. Update dependencies (if any):

   ```bash
   make deps
   ```

   Or, if you use Go modules:

   ```bash
   go mod tidy
   ```

3. Rebuild Aether:

   ```bash
   make
   ```

   Or, if you build with Go:

   ```bash
   go build ./...
   ```

If you installed Aether, you can use the command below to update your Aether version

```bash
./aether2 update
```

or add a nightly flag to update to the latest nightly version

```bash
./aether2 update --nightly
```

After updating, you can verify the version:

```bash
./aether2 --version
```

## Troubleshooting

- Ensure all prerequisites are installed and on your PATH.
- Check versions with the commands above.
- For missing dependencies, rerun the prerequisite scripts.
- For build errors, consult the output and verify your environment.

## Uninstalling

To remove Aether, delete the repository directory. To remove dependencies, use your package manager.

## Getting Help

For support, open an issue on GitHub or join the community chat (link TBD).
