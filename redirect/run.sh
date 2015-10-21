#!/bin/sh
go build main.go
./main.exe 1>>std.md 2>>err.md
