#!/bin/bash
set -eux

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  if command -v apt &>/dev/null; then
    sudo apt update
    sudo apt install -y aria2
  elif command -v dnf &>/dev/null; then
    sudo dnf install -y aria2
  elif command -v pacman &>/dev/null; then
    sudo pacman -Sy --noconfirm aria2
  else
    echo "Unsupported Linux distribution. Please install aria2 manually."
    exit 1
  fi
elif [[ "$OSTYPE" == "darwin"* ]]; then
  if command -v brew &>/dev/null; then
    brew install aria2
  else
    echo "Homebrew not found. Please install Homebrew and retry."
    exit 1
  fi
elif [[ "$OSTYPE" == "msys"* ]] || [[ "$OSTYPE" == "win32"* ]] || [[ "$OSTYPE" == "cygwin"* ]]; then
  if command -v choco &>/dev/null; then
    choco install aria2 -y
  elif command -v scoop &>/dev/null; then
    scoop install aria2
  else
    echo "Please install aria2 using Chocolatey or Scoop."
    exit 1
  fi
else
  echo "Unsupported OS. Please install aria2 manually."
  exit 1
fi 