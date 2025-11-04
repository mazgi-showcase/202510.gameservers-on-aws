#!/bin/bash

readonly PORT=8080
readonly BINARIES_BASE_DIR="$(dirname $(realpath ${BASH_SOURCE:-0}))"
readonly SERVER_BINARY_PATH="${BINARIES_BASE_DIR}/server"
readonly WRAPPER_BINARY_PATH="${BINARIES_BASE_DIR}/gameliftwrapper"

echo "Start Go Wrapper in the background and register our port"
${WRAPPER_BINARY_PATH} $PORT &
WRAPPER_PID=$!

# NOTE: For Unreal, you can run the game server binary directly but if you do run the yourgameserver.sh, make sure to make both executable!
echo "Making sure we are able to execute the server binary"
chmod +x ${SERVER_BINARY_PATH}

echo "Start game server"
${SERVER_BINARY_PATH}

echo "Game server terminated, signal wrapper so it can call ProcessEnding()"
kill -SIGINT $WRAPPER_PID
echo "Sleep for 0.3 seconds to allow the wrapper to finish"
sleep 0.3