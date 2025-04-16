#!/bin/bash

# --- Configuration ---
BIN_TARGET_DIR="/usr/local/bin"
MAN_TARGET_DIR="/usr/local/share/man/man1"
BIN_NAME="clin"
MAN_NAME="clin.1"

# --- Helper Functions ---
error_exit() {
    echo "Error: $1" >&2
    exit 1
}

check_sudo() {
    if [[ "$EUID" -ne 0 ]]; then
        echo "This script requires sudo privileges to uninstall."
        echo "Please run it with sudo: sudo $0"
        exit 1
    fi
}

remove_bin() {
    local bin_path="${BIN_TARGET_DIR}/${BIN_NAME}"
    if [[ -f "$bin_path" ]]; then
        echo "Removing binary: $bin_path"
        rm -f "$bin_path" || error_exit "Failed to remove binary."
        echo "Binary removed successfully."
    else
        echo "Binary not found: $bin_path"
    fi
}

remove_man() {
    local man_path="${MAN_TARGET_DIR}/${MAN_NAME}"
    if [[ -f "$man_path" ]]; then
        echo "Removing man page: $man_path"
        rm -f "$man_path" || error_exit "Failed to remove man page."
        echo "Man page removed successfully."
    else
        echo "Man page not found: $man_path"
    fi
}

update_man_db() {
    echo "Updating man page database..."
    if command -v mandb >/dev/null 2>&1; then
        mandb || echo "Warning: Failed to update man page database (mandb)."
    elif [[ "$(uname -s)" == "Darwin" ]]; then
        cat /etc/manpaths | while read path; do
            if [[ -n "$path" ]]; then
                man -K "$path" || echo "Warning: Failed to update man page database (man -K)."
            fi
        done
    else
        echo "Warning: Unknown system. Skipping man DB update."
    fi
    echo "Man page database update completed."
}

# --- Main Script ---

check_sudo

remove_bin
remove_man
update_man_db

echo
echo "[✓] clin has been successfully uninstalled."
echo "You may need to restart your terminal for the changes to fully apply."

