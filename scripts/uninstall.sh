#!/usr/bin/env bash
set -euo pipefail

BINARY="michi"
DEST="$HOME/.local/bin/$BINARY"

# Colors
GREEN="\033[1;32m"
YELLOW="\033[1;33m"
RED="\033[1;31m"
CYAN="\033[1;36m"
RESET="\033[0m"

echo -e "${CYAN}🚀 Starting uninstallation of $BINARY...${RESET}"

# Stop the service if running
if command -v "$DEST" &>/dev/null; then
    echo -e "${YELLOW}⏹ Stopping $BINARY service...${RESET}"
    if "$DEST" stop &>/dev/null; then
        echo -e "${GREEN}✅ $BINARY service stopped.${RESET}"
    else
        echo -e "${RED}⚠️ $BINARY stop command failed or service not running.${RESET}"
    fi
else
    echo -e "${YELLOW}⚠️ $BINARY binary not found, skipping stop.${RESET}"
fi

# Remove binary
if [ -f "$DEST" ]; then
    echo -e "${YELLOW}🗑  Removing $BINARY from $DEST...${RESET}"
    rm -f "$DEST"
    echo -e "${GREEN}✅ $BINARY removed successfully!${RESET}"
else
    echo -e "${RED}⚠️ $BINARY not found at $DEST, nothing to do.${RESET}"
    echo -e "${CYAN}💡 Perhaps it was already removed? Run the install script to reinstall.${RESET}"
    exit 0
fi

echo -e "${CYAN}ℹ️ All configuration under ~/.michi still exists. Manually remove them if needed.${RESET}"
echo -e "${GREEN}🎉 Thank you for using $BINARY!${RESET}"
