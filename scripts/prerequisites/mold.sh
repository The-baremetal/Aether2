#!/bin/bash
set -eux

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  if command -v apt &>/dev/null; then
    sudo apt update
    sudo apt install -y mold
  elif command -v dnf &>/dev/null; then
    sudo dnf install -y mold
  elif command -v pacman &>/dev/null; then
    sudo pacman -Sy --noconfirm mold
  else
    echo "Unsupported Linux distribution. Please install mold manually."
    exit 1
  fi
elif [[ "$OSTYPE" == "darwin"* ]]; then
  if command -v brew &>/dev/null; then
    brew install mold
  else
    echo "Homebrew not found. Please install Homebrew and retry."
    exit 1
  fi
elif [[ "$OSTYPE" == "msys"* ]] || [[ "$OSTYPE" == "win32"* ]] || [[ "$OSTYPE" == "cygwin"* ]]; then
  if command -v choco &>/dev/null; then
    choco install mold -y
  elif command -v scoop &>/dev/null; then
    scoop install mold
  else
    echo "Please install mold manually (see https://github.com/rui314/mold)."
    exit 1
  fi
else
  echo "Unsupported OS. Please install mold manually."
  exit 1
fi 