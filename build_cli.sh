#!/bin/bash
set -e

# Setup ENV
export GOPATH=$(pwd)
export PATH="$GOPATH/bin:$PATH"

# Binary Build
go get gopkg.in/urfave/cli.v1
go get -d ./src/...
go install ./src/cmd/cli
