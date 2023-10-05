#!/bin/bash

# Build the Dynamic Identity Server
echo "Building Dynamic Identity Server..."
go build -o dynamic_identity_server dynamic_identity_server.go
if [ $? -ne 0 ]; then
    echo "Failed to build Dynamic Identity Server."
    exit 1
fi

# Build the MFA Server
echo "Building MFA Server..."
go build -o mfa_server mfa_server.go
if [ $? -ne 0 ]; then
    echo "Failed to build MFA Server."
    exit 1
fi

# Build the API Server
echo "Building API Server..."
go build -o api_server api_server.go
if [ $? -ne 0 ]; then
    echo "Failed to build API Server."
    exit 1
fi

# Start all servers in the background

echo "Starting Dynamic Identity Server..."
./dynamic_identity_server &
DYNAMIC_SERVER_PID=$!

echo "Starting MFA Server..."
./mfa_server &
MFA_SERVER_PID=$!

echo "Starting API Server..."
./api_server &
API_SERVER_PID=$!

# Capture the PIDs so you can kill them later if needed
echo "Dynamic Identity Server PID: $DYNAMIC_SERVER_PID"
echo "MFA Server PID: $MFA_SERVER_PID"
echo "API Server PID: $API_SERVER_PID"

# Trap to handle cleanup
trap "kill $DYNAMIC_SERVER_PID $MFA_SERVER_PID $API_SERVER_PID" EXIT

# Wait indefinitely (you can kill the script to stop all servers)
wait


