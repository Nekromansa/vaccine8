#!/bin/bash
#upx vaccine8_linux.elf &&
GOOS=linux GOARCH=amd64 go build -o vaccine8_linux.elf -ldflags "-s -w" 

#&& upx vaccine8_linux.elf && mv vaccine8_linux.elf app/.
