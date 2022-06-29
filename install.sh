#!/bin/bash
set -e

INSTALL_DIR=${INSTALL_DIR:-$HOME/.local/bin}

version="1.0.0"
os="linux"

case "$OSTYPE" in
  darwin)
    os="darwin"
    ;;
  linux*)
    os="linux"
    ;;
  *)
    echo "This is an unsupported operating system."
    exit 1
    ;;
esac

unamem=$(uname -m)
if [ "$unamem" = "x86_64" ]; then
  arch="amd64"
elif [ "$unamem" = "aarch64" ]; then
  arch="arm64"
else
  echo "This is an unsupported architecture."
  exit 1
fi

echo "- Checking for local bin"
if [[ ! -d "$HOME/.local/bin" ]]; then
    echo "  - Creating local bin"
    mkdir -p "$HOME/.local/bin"
fi

echo "- Downloading binaries"
curl -L https://github.com/kzdv/cli/releases/download/v$version/release-$os-$arch.tar.gz -o /tmp/release.tar.gz

echo "- Extracting binarines"
tar -xzf /tmp/release.tar.gz -C "$INSTALL_DIR"

echo "- Cleaning up"
rm /tmp/release.tar.gz