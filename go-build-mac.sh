#!/bin/bash
#go build  -o vaccine8_mac.app -ldflags "-s -w" && mv vaccine8_mac.app app/.
#go build  -o vaccine8_mac.app -ldflags "-s -w" && upx "-9" vaccine8_mac.app && mv vaccine8_mac.app app/.
go build  -o vaccine8_mac.app -ldflags "-s -w"
