#!/bin/sh

go clean
rm -r build

# MacOS
env GOOS=darwin    GOARCH=386      go build -o build/mirrord-darwin-386
env GOOS=darwin    GOARCH=amd64    go build -o build/mirrord-darwin-amd64
env GOOS=darwin    GOARCH=arm      go build -o build/mirrord-darwin-arm
env GOOS=darwin    GOARCH=arm64    go build -o build/mirrord-darwin-arm64

# Linux
env GOOS=linux     GOARCH=386      go build -o build/mirrord-linux-386
env GOOS=linux     GOARCH=amd64    go build -o build/mirrord-linux-amd64
env GOOS=linux     GOARCH=arm      go build -o build/mirrord-linux-arm
env GOOS=linux     GOARCH=arm64    go build -o build/mirrord-linux-arm64

# Windows
env GOOS=windows   GOARCH=386      go build -o build/mirrord-windows-386.exe
env GOOS=windows   GOARCH=amd64    go build -o build/mirrord-windows-amd64.exe
env GOOS=windows   GOARCH=arm      go build -o build/mirrord-windows-arm.exe
