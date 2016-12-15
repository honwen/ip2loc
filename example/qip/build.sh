#! /bin/bash

LDFLAGS="-s -w"

if hash upx 2>/dev/null; then UPX=true; else UPX=false;fi

../../tools/gen-go-data.sh ../../17monipdb.dat main && env CGO_ENABLED=0 GOOS=linux go build -ldflags "$LDFLAGS" -o ./ip2loc
../../tools/gen-go-data.sh ../../17monipdb.dat main && env CGO_ENABLED=0 GOOS=windows go build -ldflags "$LDFLAGS" -o ./ip2loc.exe

if $UPX; then upx -9 ./ip2loc*;fi
