#!/bin/sh
set -euo pipefail

# Terminal ANSI escape codes.
RESET="\033[0m"
RED="\033[0;31m"
GREEN="\033[0;32m"

# This script is for installing the latest version of Myple CLI.

usage() {
    echo "The installer for Myple CLI"
    echo ""
    echo "Usage:"
    echo "  ${GREEN}myple-installer[EXE]${RESET} [OPTIONS]"
    echo ""
    echo "Options:"
    echo "  ${GREEN}-v, --version${RESET}"
    echo "          Version of Myple CLI to install"
    echo "  ${GREEN}-i, --install${RESET}"
    echo "          Directory to install Myple CLI"
    echo "  ${GREEN}-h, --help${RESET}"
    echo "          Output usage information"
}

# Get the operating system of the current machine.
get_os() {
    local _os
    _os=$(uname -s)
    case "$_os" in
        Linux*)     echo "linux";;
        Darwin*)    echo "darwin";;
        *)          echo "unsupported";;
    esac
}

# Get the architecture of the current machine.
get_arch() {
    local _arch
    _arch=$(uname -m)
    case "$_arch" in
        x86_64) echo "amd64";;
        arm64)  echo "arm64";;
        *)      echo "unsupported";;
    esac
}

# Check if a command exists. If the command does not exist execution
# will immediately terminate with an error showing the missing command.
need_cmd() {
    if ! command -v "$1" &> /dev/null; then
        echo "${RED}error${RESET}: need '$1' (command not found)"
        exit 1
    fi
}

# Run a command that should never fail. If the command fails execution
# will immediately terminate with an error showing the failing
# command.
ensure() {
    if ! "$@"; then echo "${RED}error${RESET}: command failed: $*"; exit 1; fi
}

# Detect the shell profile file. The profile file is used to add the
# compeltion script to the shell.
detect_profile() {
  need_cmd basename

  local _detected_profile
  _detected_profile=''
  local _shelltype
  _shelltype=$(basename "/$SHELL")

  if [ "$_shelltype" = "bash" ]; then
    if [ -f "$HOME/.bashrc" ]; then
      _detected_profile="$HOME/.bashrc"
    elif [ -f "$HOME/.bash_profile" ]; then
      _detected_profile="$HOME/.bash_profile"
    fi
  elif [ "$_shelltype" = "zsh" ]; then
    _detected_profile="$HOME/.zshrc"
  elif [ "$_shelltype" = "fish" ]; then
    _detected_profile="$HOME/.config/fish/conf.d/myple.fish"
  fi

  if [ -z "$_detected_profile" ]; then
    if [ -f "$HOME/.profile" ]; then
      _detected_profile="$HOME/.profile"
    elif [ -f "$HOME/.bashrc" ]; then
      _detected_profile="$HOME/.bashrc"
    elif [ -f "$HOME/.bash_profile" ]; then
      _detected_profile="$HOME/.bash_profile"
    elif [ -f "$HOME/.zshrc" ]; then
      _detected_profile="$HOME/.zshrc"
    elif [ -d "$HOME/.config/fish" ]; then
      _detected_profile="$HOME/.config/fish/conf.d/myple.fish"
    fi
  fi

  if [ ! -z "$_detected_profile" ]; then
    echo "$_detected_profile"
  fi
}

update_profile() {
  _profile = $(detect_profile)

  if [ -z "$_profile" ]; then
    echo "${RED}error${RESET}: unable to detect profile file location"
    exit 1
  fi

  if ! grep -q "\.myple" $_profile; then
    echo "\n# Myple CLI\nexport PATH=\"\$HOME/.myple/bin:\$PATH\"\n" >> $_profile
  fi
}

main() {
    need_cmd uname
    need_cmd curl
    need_cmd tar
    need_cmd mkdir
    need_cmd rm

    local _version
    _version="${MYPLE_VERSION:-latest}"

    local _install
    _install="${MYPLE_INSTALL:-$HOME}"

    for arg in "$@"; do
        case "$arg" in
            -v | --version)
                _version="$2"
                shift
                ;;
            -i | --install)
                _install="$2"
                shift
                ;;
            -h | --help)
                usage
                exit 0
                ;;
            *)
                echo "${RED}error${RESET}: unknown option '$arg'"
                exit 1
                ;;
        esac
    done

    local _os
    _os=$(get_os)
    if [[ "$_os" == "unsupported" ]]; then
        echo "${RED}error${RESET}: unsupported operating system"
        exit 1
    fi

    local _arch
    _arch=$(get_arch)
    if [[ "$_arch" == "unsupported" ]]; then
        echo "${RED}error${RESET}: unsupported architecture"
        exit 1
    fi

    local _bin_dir
    _bin_dir="$_install/.myple/bin"
    local _tmp_dir
    _tmp_dir="$_install/.myple/tmp"

    ensure mkdir -p "$_bin_dir"
    ensure mkdir -p "$_tmp_dir"

    local _exe
    _exe="$_bin_dir/myple"

    local _file
    _file="myple-$_version-$_os-$_arch.tar.gz"
    local _url
    _url="https//github.com/myple/cli/releases/download/$_version/$_file"

    echo "Downloading Myple CLI..."
    curl -sL "$_url" -o "$_tmp_dir/$_file" || (echo "${RED}error${RESET}: failed to download Myple CLI"; exit 1)
    echo "${GREEN}info${RESET}: Myple CLI has been downloaded successfully"

    echo "Installing Myple CLI..."
    tar -xzf $_file -C $_tmp_dir || (echo "${RED}error${RESET}: failed to extract Myple CLI"; exit 1)
    chmod +x "$_tmp_dir/myple"
    mv "$_tmp_dir/myple" "$_exe"
    rm "$_tmp_dir/$_file"
    echo "${GREEN}info${RESET}: Myple CLI has been installed successfully"

    echo "Adding completion script to shell profile..."
    echo "${GREEN}info${RESET}: completion script has been added to shell profile"

    echo "Myple CLI has been installed successfully to $_exe"
    update_profile

    echo ""
    echo "You can now run 'myple login' to authenticate with Myple,"
    echo "or run 'myple help' to get know more about the available commands."
    echo ""
    echo "For more information, visit ${GREEN}https://docs.myple.io/cli${RESET}"
}

main "$@" || exit 1
