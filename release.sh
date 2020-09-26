#!/usr/bin/sh

pkgname=iwsp

GOOS=linux go build -ldflags "-w -s -X main.version=$1" -o release/${pkgname}-linux .
GOOS=darwin go build -ldflags "-w -s -X main.version=$1" -o release/${pkgname}-macos .
GOOS=windows go build -ldflags "-w -s -X main.version=$1" -o release/${pkgname}-windows.exe .
