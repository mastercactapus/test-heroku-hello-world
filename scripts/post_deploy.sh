#!/usr/bin/env bash
go build -o hello
./hello -set-pub=$(heroku info -s | grep web-url | cut -d= -f2)