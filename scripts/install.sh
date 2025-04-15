#!/bin/bash

# --- Configuration ---
REPO_OWNER="YOUR_GITHUB_USERNAME"       # Replace with your GitHub username
REPO_NAME="YOUR_GITHUB_REPO"           # Replace with your GitHub repository name
BIN_RELEASE_ASSET_LINUX_X64="clin-linux-x64"     # Replace with your Linux x64 binary asset name
BIN_RELEASE_ASSET_MACOS_X64="clin-macos-x64"     # Replace with your macOS Intel x64 binary asset name
BIN_RELEASE_ASSET_MACOS_ARM64="clin-macos-arm64"   # Replace with your macOS ARM64 binary asset name
MAN_RELEASE_ASSET="clin.1"            # Replace with the name of your man page asset in the release
BIN_TARGET_DIR="/usr/local/bin"
MAN_TARGET_DIR="/usr/local/share/man/man1"
VERSION="1.0.0" #  Set your initial version here

# --- Helper Functions ---
error_exit() {
    echo "Error: $1" >&2
    exit 1
}

check_sudo() {
    if [[ "$EUID" -ne 0 ]]; then
        echo "This script requires sudo privileges to install."
        echo "Please run it with sudo: sudo $0"
        exit 1
    fi
}

download_asset() {
    local asset_name="$1"
    local download_url="https://github.com/${REPO_OWNER}/${REPO_NAME}/releases/latest/download/${asset_name}"
    echo "Downloading ${asset_name}..."
    curl -L -o "${asset_name}" "${download_url}" || error_exit "Failed to download ${asset_name}."
    echo "Downloaded ${asset_name} successfully."
}

install_bin() {
    echo "Installing binary to ${BIN_TARGET_DIR}/clin..."
    mv "${BIN_RELEASE_ASSET}" "${BIN_TARGET_DIR}/clin" || error_exit "Failed to move and rename binary."
    chmod +x "${BIN_TARGET_DIR}/clin" || error_exit "Failed to make binary executable."
    echo "Binary installed successfully."
}

install_man() {
    echo "Installing man page to ${MAN_TARGET_DIR}/${MAN_RELEASE_ASSET}..."
    mkdir -p "${MAN_TARGET_DIR}" || error_exit "Failed to create man page directory."
    mv "${MAN_RELEASE_ASSET}" "${MAN_TARGET_DIR}/" || error_exit "Failed to move man page."
    echo "Man page installed successfully."
}

update_man_db() {
    echo "Updating man page database..."
    if command -v mandb >/dev/null 2>&1; then
        sudo mandb || echo "Warning: Failed to update man page database (mandb)."
    elif [[ "$(uname -s)" == "Darwin" ]]; then
        sudo cat /etc/manpaths | while read path; do
            if [[ -n "$path" ]]; then
                sudo man -K "$path" || echo "Warning: Failed to update man page database (man -K)."
            fi
        done
    else
        echo "Warning: Could not find 'mandb' or macOS man update mechanism. Man page might not be immediately available."
    fi
    echo "Man page database update initiated."
}

# --- Main Script ---

check_sudo

os=$(uname -s)
arch=$(uname -m)

echo "Detected OS: $os"
echo "Detected Architecture: $arch"

case "$os" in
    Linux*)
        if [[ "$arch" == "x86_64" ]]; then
            BIN_RELEASE_ASSET="$BIN_RELEASE_ASSET_LINUX_X64"
            download_asset "$BIN_RELEASE_ASSET"
        else
            error_exit "Unsupported Linux architecture: $arch"
        fi
        ;;
    Darwin*)
        if [[ "$arch" == "x86_64" ]]; then
            BIN_RELEASE_ASSET="$BIN_RELEASE_ASSET_MACOS_X64"
            download_asset "$BIN_RELEASE_ASSET"
        elif [[ "$arch" == "arm64" ]]; then
            BIN_RELEASE_ASSET="$BIN_RELEASE_ASSET_MACOS_ARM64"
            download_asset "$BIN_RELEASE_ASSET"
        else
            error_exit "Unsupported macOS architecture: $arch"
        fi
        ;;
    *)
        error_exit "Unsupported operating system: $os"
        ;;
esac

download_asset "${MAN_RELEASE_ASSET}"

install_bin
install_man
update_man_db

echo "clin has been successfully installed. You might need to open a new terminal session."
echo "clin version: $VERSION" #show version

exit 0

