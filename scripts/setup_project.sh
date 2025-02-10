#!/bin/bash

# Check if the correct number of arguments are provided
if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <project_name> <function_name>"
    exit 1
fi

# Assign arguments to variables
PROJECT_NAME=$1
FUNCTION_NAME=$2

# Handle case where the project directory already exists
if [ -d "$PROJECT_NAME" ]; then
    echo "Project directory already exists"
    exit 1
fi

# Create project directory
mkdir -p $PROJECT_NAME

# Create main.go file
cat <<EOL > $PROJECT_NAME/main.go
package main

import "fmt"

func main() {
    $FUNCTION_NAME()
}

func $FUNCTION_NAME() {
    fmt.Println("Hello from $FUNCTION_NAME")
}
EOL

# Open the file with VS Code
code $PROJECT_NAME/main.go

# Run the project
./run_go.sh $PROJECT_NAME

# Show the run command for the user
echo "To run the project, use the following command:"
echo ""
COMMAND="./run_go.sh $PROJECT_NAME"
echo "$COMMAND"

# Copy that command to their clipboard
echo "$COMMAND" | pbcopy
