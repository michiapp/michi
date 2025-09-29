#!/usr/bin/env bash
set -euo pipefail

REPO="OrbitalJin/michi"
BINARY="michi"

# Colors
GREEN="\033[1;32m"
YELLOW="\033[1;33m"
RED="\033[1;31m"
CYAN="\033[1;36m"
RESET="\033[0m"

echo -e "${CYAN}🚀 Starting installation of $BINARY...${RESET}"

# Detect OS
OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"

case "$ARCH" in
    x86_64) ARCH="amd64" ;;
    aarch64 | arm64) ARCH="arm64" ;;
    *) echo -e "${RED}❌ Unsupported architecture: $ARCH${RESET}"; exit 1 ;;
esac

VERSION="${VERSION:-latest}"

# Fetch latest version if needed
if [ "$VERSION" = "latest" ]; then
    echo -e "${CYAN}🔍 Fetching latest release version...${RESET}"
    VERSION=$(curl -fsSL "https://api.github.com/repos/$REPO/releases/latest" \
        | grep -oP '"tag_name": "\K(.*)(?=")')
fi

echo -e "${CYAN}⬇️ Installing $BINARY $VERSION for $OS/$ARCH...${RESET}"

# Build download URL
TARFILE="${BINARY}_${OS}_${ARCH}.tar.gz"
URL="https://github.com/$REPO/releases/download/$VERSION/$TARFILE"
echo -e "${YELLOW}🌐 Download URL: $URL${RESET}"

TMPDIR=$(mktemp -d)
curl -fsSL "$URL" -o "$TMPDIR/$TARFILE"

# Extract
echo -e "${CYAN}📦 Extracting archive...${RESET}"
tar -xzf "$TMPDIR/$TARFILE" -C "$TMPDIR"

# Install
DEST="$HOME/.local/bin/$BINARY"
mkdir -p "$(dirname "$DEST")"

mv "$TMPDIR/$BINARY" "$DEST"
chmod +x "$DEST"

echo -e "${GREEN}✅ Installed $BINARY to $DEST${RESET}"
echo -e "${CYAN}💡 Make sure $HOME/.local/bin is in your PATH.${RESET}"
echo -e "${GREEN}🎉 Installation complete!${RESET}"
echo -e "${GREEN}🌟 Enjoy using $BINARY!${RESET}"
echo -e "${CYAN}🚀 Get started by running: $BINARY --help${RESET}"
