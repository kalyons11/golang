#!/bin/bash

run_go() {
  clear

  if [ -z "$1" ]; then
    echo "Error: No source file provided."
    exit 1
  fi

  SOURCE_DIR="$1"
  # Trim any slashes from the end of the source directory
  SOURCE_DIR="${SOURCE_DIR%/}"

  pushd "$SOURCE_DIR" || { echo "Error: Failed to change directory to '$SOURCE_DIR'."; exit 1; }

  # Run the program, forwarding any additional command line arguments to the executable
  go mod init kevin/$SOURCE_DIR
  go run .

  # Return to the original directory
  popd || { echo "Error: Failed to return to the original directory."; exit 1; }
}

run_go "$@"
