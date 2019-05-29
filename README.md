# Hello World

Test app for heroku

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

## Building

```bash
go build -o hello
```

## Setup

```bash
./hello -db-url=postgres://localhost:5432/test_db?sslmode=disable -set-pub=https://myapp.example.com
```

## Run

```bash
./hello  -db-url=postgres://localhost:5432/test_db?sslmode=disable
```

It will be running at http://localhost:8080
