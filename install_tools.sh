#!/bin/zsh
# install_tools.sh - Install subfinder, assetfinder, and findomain for 3subs

set -e

echo "[+] Installing subfinder..."
GO111MODULE=on go install -v github.com/projectdiscovery/subfinder/v2/cmd/subfinder@latest

echo "[+] Installing assetfinder..."
GO111MODULE=on go install -v github.com/tomnomnom/assetfinder@latest

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
if [[ "$OS" == "darwin" ]]; then
  echo "[+] Installing findomain via Homebrew (macOS)..."
  if ! command -v brew >/dev/null 2>&1; then
    echo "[!] Homebrew is not installed. Please install Homebrew first: https://brew.sh/"
    exit 1
  fi
  brew install findomain
else
  echo "[+] Installing findomain (Linux)..."
  FINDOMAIN_VERSION=$(curl -s https://api.github.com/repos/findomain/findomain/releases/latest | grep 'tag_name' | cut -d'\"' -f4)
  ARCH=$(uname -m)
  if [[ "$ARCH" == "x86_64" ]]; then ARCH="amd64"; fi
  if [[ "$ARCH" == "aarch64" ]]; then ARCH="arm64"; fi
  FILENAME="findomain-$OS-$ARCH.zip"
  rm -f $FILENAME findomain # Clean up any previous files
  curl -LO https://github.com/findomain/findomain/releases/download/$FINDOMAIN_VERSION/$FILENAME
  if unzip -o $FILENAME; then
    chmod +x findomain
    sudo mv findomain /usr/local/bin/
    rm $FILENAME
    echo "[✓] findomain installed."
  else
    echo "[!] Failed to unzip $FILENAME. Please check your architecture and the downloaded file."
    exit 1
  fi
fi

echo "[✓] All tools installed. Make sure $HOME/go/bin is in your PATH."
