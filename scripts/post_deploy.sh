#!/usr/bin/env bash
go build -o hello
./hello -db-url=postgres://localhost:5432/test_db?sslmode=disable -set-pub=$(heroku info -s | grep web-url | cut -d= -f2)