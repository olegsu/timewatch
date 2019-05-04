#!/bin/bash
set -e
OUTFILE=/usr/local/bin/timewatchdev
go build -o $OUTFILE main.go

chmod +x $OUTFILE