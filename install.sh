#!/usr/bin/env bash
TMP_DIR="/tmp/goft-installer-"$(echo $RANDOM | shasum | head -c 20)

mkdir -p "$TMP_DIR"

# Detect OS and ARCH
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

echo "Your OS is $OS and ARCH is $ARCH"

if [ "$OS" != "darwin" ]; then
  echo "Currently this script only works with MacOS :("
  exit 1
fi

echo "Downloading goft binary..."

URL_DOWNLOAD="https://github.com/mehdibo/goft/releases/latest/download/goft-$OS"

if [ "$ARCH" == "x86_64" ]; then
  URL_DOWNLOAD=$URL_DOWNLOAD"-amd64"
fi

curl -s -L -o "$TMP_DIR/goft" "$URL_DOWNLOAD"
chmod u+x "$TMP_DIR/goft"

echo "Moving goft binary to /usr/local/bin, this will be run as root"
sudo mv "$TMP_DIR/goft" /usr/local/bin

# Download config sample if not existing
if [ -f "$HOME/.goft.yml" ]; then
    echo "Config file already exists, skipping..."
else
    echo "Downloading config file"
    curl -s -L -o "$TMP_DIR/config.yml" https://raw.githubusercontent.com/mehdibo/goft/main/config.example.yml
    mv "$TMP_DIR/config.yml" "$HOME/.goft.yml"
    echo "Don't forget to update $HOME/.goft.yml with your credentials"
fi

echo "Try running: goft --help"

rm -rf "$TMP_DIR"