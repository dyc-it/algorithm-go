#!/usr/bin/env bash

cd /Users/dyc/Documents/code/go-project/src

# format the whole directory
go fmt algorithm-go/...

# run all unit tests in this directory
go test algorithm-go/...
