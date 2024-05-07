#!/bin/sh
set -euo pipefail

# Terminal ANSI escape codes.
RESET="\033[0m"
RED="\033[0;31m"
GREEN="\033[0;32m"
YELLOW="\033[0;33m"
BLUE="\033[0;34m"

# This script is for uninstalling Myple CLI.

# Check if a command exists. If the command does not exist execution
# will immediately terminate with an error showing the missing command.
need_cmd() {
    if ! command -v "$1" &> /dev/null; then
        echo "${RED}error${RESET}: need '$1' (command not found)"
        exit 1
    fi
}

# Detect the shell profile file. The profile file is used to add the
# compeltion script to the shell.
detect_profile() {
  need_cmd basename

  local _detected_profile
  _detected_profile=''
  local _shell
  _shell=$(basename "/$SHELL")

  if [ "$_shell" = "bash" ]; then
    if [ -f "$HOME/.bashrc" ]; then
      _detected_profile="$HOME/.bashrc"
    elif [ -f "$HOME/.bash_profile" ]; then
      _detected_profile="$HOME/.bash_profile"
    fi
  elif [ "$_shell" = "zsh" ]; then
    _detected_profile="$HOME/.zshrc"
  elif [ "$_shell" = "fish" ]; then
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
  need_cmd basename
  need_cmd rm
  need_cmd sed

  local _profile
  _profile=$(detect_profile)

  if [ -z "$_profile" ]; then
    echo "${YELLOW}warning${RESET}: profile not found"
    exit 1
  fi

  if [ "$_profile" = "$HOME/.config/fish/conf.d/myple.fish" ]; then
    rm -f "$_profile"
  elif [ -f "$_profile" ]; then
    sed -i '/\n# Myple CLI\nexport PATH=\"\$HOME/.myple/bin:\$PATH\"\n/d' "$_profile"
    # TODO remove the completion script from the zsh and bash profiles
  fi
}

main() {
  echo "Uninstalling Myple CLI..."
  need_cmd rm
  rm -rf "$HOME/.myple"
   echo "${GREEN}info${RESET}: Myple CLI has been uninstalled"

  echo "Removing completion script..."
  update_profile
  echo "${GREEN}info${RESET}: completion script has been removed"
}

main "$@" || exit 1
