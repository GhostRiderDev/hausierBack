#!/bin/bash

export GOOS=windows
rm -rf build
mkdir build
go build -o build/app.exe cmd/main.go
build/app.exe

