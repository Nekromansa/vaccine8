#!/bin/bash
GOOS=windows GOARCH=386 go build  -o vaccine8_win.exe -ldflags "-s -w" && upx vaccine8_win.exe && mv vaccine8_win.exe app/.
