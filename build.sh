#!/bin/bash
echo Downloading dependencies...
go mod download
go mod verify
go build -v -o ./dist/go-short-app