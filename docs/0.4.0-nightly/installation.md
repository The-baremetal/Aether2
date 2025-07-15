# Aether Installation Guide

## Introduction

Aether is a modern programming language and toolchain. This guide provides detailed, step-by-step instructions for installing prerequisites, downloading binaries, building Aether from source, and updating or uninstalling on various operating systems and distributions.

---

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [Downloading Binaries](#downloading-binaries)
3. [Building from Source](#building-from-source)
4. [Platform-Specific Installation Guides](#platform-specific-installation-guides)

   * [Linux](#linux)
   * [macOS](#macos)
   * [Windows](#windows)
5. [Updating Aether](#updating-aether)
6. [Troubleshooting](#troubleshooting)
7. [Uninstalling](#uninstalling)
8. [Getting Help](#getting-help)

---

## Prerequisites

Before installing or building Aether, ensure the following tools are installed and available on your system PATH:

* **Go** (>=1.20)
* **LLVM** (>=15)
* **aria2**
* **mold** linker

### Checking and Installing Prerequisites

#### Go

* **Check**:  `go version`
* **Install**:

  * **Debian/Ubuntu**: `sudo apt install golang`
  * **Fedora/RHEL**: `sudo dnf install golang`
  * **Arch Linux**: `sudo pacman -Sy go`
  * **macOS**: `brew install go`
  * **Windows**: Download from [golang.org](https://golang.org/dl/) and follow the installer.

#### LLVM

* **Check**: `llvm-config --version`
* **Install**:

  * **Debian/Ubuntu**: `sudo apt install llvm`
  * **Fedora/RHEL**: `sudo dnf install llvm`
  * **Arch Linux**: `sudo pacman -Sy llvm`
  * **macOS**: `brew install llvm`
  * **Windows**: Download from [LLVM Releases](https://releases.llvm.org/) and follow installer.

#### aria2

* **Check**: `aria2c --version`
* **Install**:

  * **Linux**: `sudo apt install aria2` / `sudo dnf install aria2` / `sudo pacman -Sy aria2`
  * **macOS**: `brew install aria2`
  * **Windows (Chocolatey)**: `choco install aria2` or **(Scoop)**: `scoop install aria2`

#### mold linker

* **Check**: `mold --version`
* **Install**:

  * **Linux**: clone and build from source or use distro package if available
  * **macOS**: `brew install mold`
  * **Windows**: build via MSYS2 or WSL

---

## Downloading Binaries {#downloading-binaries}

Pre-built Aether binaries are available for various platforms. Use **aria2** to download rapidly:

```bash
aria2c -x16 -s16 -j1 "$(curl -s https://api.github.com/repos/The-baremetal/Aether2/releases/latest | grep browser_download_url | grep aether-linux_amd64.tar | cut -d '"' -f4)"
```

To download multiple assets in parallel:

```bash
aria2c -x16 -s16 -j4 \
  "$(curl -s https://api.github.com/repos/The-baremetal/Aether2/releases/latest | grep browser_download_url | grep aether-linux_amd64.tar | cut -d '"' -f4)" \
  "$(curl -s https://api.github.com/repos/The-baremetal/Aether2/releases/latest | grep browser_download_url | grep aether-linux_amd64.deb | cut -d '"' -f4)" \
  "$(curl -s https://api.github.com/repos/The-baremetal/Aether2/releases/latest | grep browser_download_url | grep aether-linux_amd64.rpm | cut -d '"' -f4)"
```

---

## Building from Source {#building-from-source}

1. **Clone the repository**:

   ```bash
   git clone https://github.com/The-baremetal/Aether2.git
   cd Aether2
   ```

2. **Install prerequisites** (see above).

3. **Build**:

   * **Make**:

     ```bash
     make build
     ```

   * **Go**:

     ```bash
     go build ./...
     ```

4. **Run tests**:

   ```bash
   go test ./...
   ```

5. **Verify CLI**:

   ```bash
   ./build/bin/aether2 --version
   ```

---

## Platform-Specific Installation Guides {#platform-specific-installation-guides}

### Linux {#linux}

#### Debian / Ubuntu

```bash
sudo apt update
sudo apt install golang llvm aria2 mold
```

```bash
aria2c -x16 -s16 -j1 "$(curl -s https://api.github.com/repos/The-baremetal/Aether2/releases/latest | grep browser_download_url | grep aether-linux_amd64.deb | cut -d '"' -f4)" && sudo dpkg -i aether-linux_amd64.deb && aether --version
```

#### Fedora / RHEL

```bash
sudo dnf install golang llvm aria2 mold
```

```bash
aria2c -x16 -s16 -j1 "$(curl -s https://api.github.com/repos/The-baremetal/Aether2/releases/latest | grep browser_download_url | grep aether-linux_amd64.rpm | cut -d '"' -f4)" && sudo rpm -i aether-linux_amd64.rpm && aether --version
```

#### Arch Linux

```bash
sudo pacman -Sy go llvm aria2 mold
```

```bash
aria2c -x16 -s16 -j1 "$(curl -s https://api.github.com/repos/The-baremetal/Aether2/releases/latest | grep browser_download_url | grep aether-linux_amd64.tar | cut -d '"' -f4)"
tar -xf aether-linux_amd64.tar
sudo mv aether /usr/local/bin/
aether --version
```

#### Other Distros (Generic TAR)

```bash
aria2c -x16 -s16 -j1 "$(curl -s https://api.github.com/repos/The-baremetal/Aether2/releases/latest | grep browser_download_url | grep aether-linux_amd64.tar | cut -d '"' -f4)"
tar -xf aether-linux_amd64.tar
chmod +x aether
sudo mv aether /usr/local/bin/
aether --version
```

### macOS {#macos}

1. Install with Homebrew:

   ```bash
   brew install go llvm aria2 mold
   ```

2. Download:

   ```bash
   aria2c -x16 -s16 -j1 "$(curl -s https://api.github.com/repos/The-baremetal/Aether2/releases/latest | grep browser_download_url | grep aether-darwin_amd64.tar | cut -d '"' -f4)"
   ```

3. Install:

   ```bash
   tar -xf aether-darwin_amd64.tar
   chmod +x aether
   sudo mv aether /usr/local/bin/
   aether --version
   ```

### Windows {#windows}

#### Recommended: Tarball Install

The preferred way to install Aether on Windows is to use the tarball release. This preserves the correct AETHERROOT layout for the standard library and binary.

**Steps:**

1. **Download the latest tarball** for Windows from the [releases page](https://github.com/The-baremetal/Aether2/releases).
2. **Extract the tarball** to a directory, e.g. `C:\Program Files\Aether` (you may need admin rights).
3. **Set the AETHERROOT environment variable:**
   - **Command Prompt:**

     ```cmd
     set AETHERROOT=C:\Program Files\Aether
     ```

   - **PowerShell:**

     ```powershell
     $env:AETHERROOT = "C:\Program Files\Aether"
     ```

4. (Optional) Add `C:\Program Files\Aether\bin` to your `PATH` for easy access to the `aether` command.

After extraction, your directory should look like:

```
C:\Program Files\Aether\
  bin\
    aether.exe
  packages\
    fmt\
    math\
    string\
    ...
```

This ensures the Aether compiler can always find the standard library in the correct place.

For more details on the install layout and standard library, see [AETHERROOT Install Layout](#aetherroot-install-layout) and [stdlib.md](./stdlib.md).

---

## Updating Aether {#updating-aether}

To update your installation:

```bash
cd Aether2
git pull origin main
make deps   # if using Make
go mod tidy # if using Go modules
make         # or go build ./...
```

Or use the built-in updater:

```bash
./aether2 update           # latest stable
./aether2 update --nightly # latest nightly
./aether2 --version
```

---

## Troubleshooting {#troubleshooting}

* Verify prerequisites on `PATH`.
* Re-run prerequisite scripts if missing.
* Inspect error messages and environment variables.

---

## Uninstalling {#uninstalling}

* **Binaries**: Remove `/usr/local/bin/aether` or the installed package via your package manager.
* **Source**: Delete the `Aether2` directory.

---

## Getting Help {#getting-help}

For support, open an issue on [GitHub](https://github.com/The-baremetal/Aether2/issues) or join the community chat.

---

## AETHERROOT Install Layout

After installation (via tarball, .deb, .rpm, .pkg, etc.), Aether will be installed in the following structure:

```
/usr/local/aether/
  bin/
    aether         # The Aether compiler binary
  packages/
    fmt/
    math/
    string/
    ...            # (copied directly from the source /packages)
```

- On Windows, this will be under `C:\Program Files\Aether\`.
- The installer also creates a symlink at `/usr/local/bin/aether` for convenience (on Unix-like systems).

**AETHERROOT** should be set to `/usr/local/aether` (or the equivalent install path on your system).

### How to Set AETHERROOT
- **Unix/Linux/macOS:**
  ```sh
  export AETHERROOT=/usr/local/aether
  ```
- **Windows (cmd):**
  ```cmd
  set AETHERROOT=C:\Program Files\Aether
  ```
- **Windows (PowerShell):**
  ```powershell
  $env:AETHERROOT = "C:\Program Files\Aether"
  ```

For more details on the standard library and import resolution, see [stdlib.md](./stdlib.md).

---
