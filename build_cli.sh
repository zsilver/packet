#!/bin/bash
set -e

# Setup ENV
export GOPATH=$(pwd)
export PATH="$GOPATH/bin:$PATH"

# Binary Build
go get ./src/...
go install ./src/cmd/cli
