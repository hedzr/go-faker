# hedzr/go-faker

[![Go](https://github.com/hedzr/go-faker/actions/workflows/go.yml/badge.svg)](https://github.com/hedzr/go-faker/actions/workflows/go.yml)

A template repository to build your first app based [cmdr](https://github.com/hedzr/cmdr).

powered by [cmdr](https://github.com/hedzr/cmdr).

## Features

## Getting Started

To run the CLI app:

```bash
go generate ./...          # run it once at least, for gen the swagger-doc files from skeletons
go run ./cli/app/cli/app   # build the mainly main.go
```

### Use Makefile for building and CI

You may use `make` simply:

```bash
make help    # list all available make targets, such as info, build, ...
make info    # print and review the golang build env

make build
```

## LICENSE

MIT


