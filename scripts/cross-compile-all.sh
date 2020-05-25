#!/bin/sh

go clean
rm -r build

# From `go tool dist list`.
env GOOS=aix	   GOARCH=ppc64	   go build -o build/mirrord-aix-ppc64
env GOOS=android   GOARCH=386      go build -o build/mirrord-android-386
env GOOS=android   GOARCH=amd64    go build -o build/mirrord-android-amd64
env GOOS=android   GOARCH=arm      go build -o build/mirrord-android-arm
env GOOS=android   GOARCH=arm64    go build -o build/mirrord-android-arm64
env GOOS=darwin    GOARCH=386      go build -o build/mirrord-darwin-386
env GOOS=darwin    GOARCH=amd64    go build -o build/mirrord-darwin-amd64
env GOOS=darwin    GOARCH=arm      go build -o build/mirrord-darwin-arm
env GOOS=darwin    GOARCH=arm64    go build -o build/mirrord-darwin-arm64
env GOOS=dragonfly GOARCH=amd64    go build -o build/mirrord-dragonfly-amd64
env GOOS=freebsd   GOARCH=386      go build -o build/mirrord-freebsd-386
env GOOS=freebsd   GOARCH=amd64    go build -o build/mirrord-freebsd-amd64
env GOOS=freebsd   GOARCH=arm      go build -o build/mirrord-freebsd-arm
env GOOS=illumos   GOARCH=amd64    go build -o build/mirrord-illumos-amd64
env GOOS=js        GOARCH=wasm     go build -o build/mirrord-js-wasm
env GOOS=linux     GOARCH=386      go build -o build/mirrord-linux-386
env GOOS=linux     GOARCH=amd64    go build -o build/mirrord-linux-amd64
env GOOS=linux     GOARCH=arm      go build -o build/mirrord-linux-arm
env GOOS=linux     GOARCH=arm64    go build -o build/mirrord-linux-arm64
env GOOS=linux     GOARCH=mips     go build -o build/mirrord-linux-mips
env GOOS=linux     GOARCH=mips64   go build -o build/mirrord-linux-mips64
env GOOS=linux     GOARCH=mips64le go build -o build/mirrord-linux-mips64le
env GOOS=linux     GOARCH=mipsle   go build -o build/mirrord-linux-mipsle
env GOOS=linux     GOARCH=ppc64    go build -o build/mirrord-linux-ppc64
env GOOS=linux     GOARCH=ppc64le  go build -o build/mirrord-linux-ppc64le
env GOOS=linux     GOARCH=s390x    go build -o build/mirrord-linux-s390x
env GOOS=nacl      GOARCH=386      go build -o build/mirrord-nacl-386
env GOOS=nacl      GOARCH=amd64p32 go build -o build/mirrord-nacl-amd64p32
env GOOS=nacl      GOARCH=arm      go build -o build/mirrord-nacl-arm
env GOOS=netbsd    GOARCH=386      go build -o build/mirrord-netbsd-386
env GOOS=netbsd    GOARCH=amd64    go build -o build/mirrord-netbsd-amd64
env GOOS=netbsd    GOARCH=arm      go build -o build/mirrord-netbsd-arm
env GOOS=netbsd    GOARCH=arm64    go build -o build/mirrord-netbsd-arm64
env GOOS=openbsd   GOARCH=386      go build -o build/mirrord-openbsd-386
env GOOS=openbsd   GOARCH=amd64    go build -o build/mirrord-openbsd-amd64
env GOOS=openbsd   GOARCH=arm      go build -o build/mirrord-openbsd-arm
env GOOS=openbsd   GOARCH=arm64    go build -o build/mirrord-openbsd-arm64
env GOOS=plan9     GOARCH=386      go build -o build/mirrord-plan9-386
env GOOS=plan9     GOARCH=amd64    go build -o build/mirrord-plan9-amd64
env GOOS=plan9     GOARCH=arm      go build -o build/mirrord-plan9-arm
env GOOS=solaris   GOARCH=amd64    go build -o build/mirrord-solaris-amd64
env GOOS=windows   GOARCH=386      go build -o build/mirrord-windows-386.exe
env GOOS=windows   GOARCH=amd64    go build -o build/mirrord-windows-amd64.exe
env GOOS=windows   GOARCH=arm      go build -o build/mirrord-windows-arm.exe
