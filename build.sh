#!/bin/sh
go mod download
go build -v -o ./out/passgengo .