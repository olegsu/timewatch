#!/bin/bash
set -e
OUTFILE=/usr/local/bin/timewatch
go build -o $OUTFILE main.go

chmod +x $OUTFILE