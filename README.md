# hedzr/go-faker

[![Go](https://github.com/hedzr/go-faker/actions/workflows/go.yml/badge.svg)](https://github.com/hedzr/go-faker/actions/workflows/go.yml)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/hedzr/go-faker.svg?label=release)](https://github.com/hedzr/go-faker/releases)
[![](https://img.shields.io/badge/go-dev-green)](https://pkg.go.dev/github.com/hedzr/go-faker)
| ![Docker Pulls](https://img.shields.io/docker/pulls/hedzr/faker)
![Docker Image Version (latest semver)](https://img.shields.io/docker/v/hedzr/faker)
![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/hedzr/faker)

<!-- [![GitHub tag](https://img.shields.io/github/tag/hedzr/consul-tags.svg)]() -->
<!-- [![ImageLayers Size](https://img.shields.io/imagelayers/image-size/hedzr/consul-tags/latest.svg)]() -->

<!-- [![GitHub version](https://badge.fury.io/gh/hedzr%2Fconsul-tags.svg)](https://badge.fury.io/gh/hedzr%2Fconsul-tags)
-->
<!--
[![license](https://img.shields.io/github/license/hedzr/go-faker.svg)](https://pkg.go.dev/github.com/hedzr/go-faker)
[![go.dev](https://img.shields.io/badge/go.dev-reference-green)](https://pkg.go.dev/github.com/hedzr/go-faker)
[![Go Report Card](https://goreportcard.com/badge/github.com/hedzr/go-faker)](https://goreportcard.com/report/github.com/hedzr/go-faker)
[![codecov](https://codecov.io/gh/hedzr/go-faker/branch/master/graph/badge.svg)](https://codecov.io/gh/hedzr/go-faker)
[![Coverage Status](https://coveralls.io/repos/github/hedzr/go-faker/badge.svg?branch=master)](https://coveralls.io/github/hedzr/go-faker?branch=master)
-->

A simple CLI app based [cmdr](https://github.com/hedzr/cmdr) and golang.

> powered by [cmdr](https://github.com/hedzr/cmdr).

## Features

`faker` will generate the faked records. It bases on [dmgk/faker](https://github.com/dmgk/faker).

The avaliable commands are:

```bash
Commands:
  a, addr, address                          generate address
  app                                       generate app names
  avatar                                    generate Avatar names
  b, bitcoin, btc                           generate Bitcoin names
  bz, business, credit-card                 generate Bitcoin names
  c, code, codes                            generate codes (ISBN10, ISBN13, EAN13, EAN8, RUT, ABN)
  cc, commerce                              generate Commerce names
  corp, company                             generate Company names
  f, finance                                generate Finance names (visa, mastercard, ...)
  hh, hacker, hack                          generate Hacker names
  hp, hacker-phrases                        generate Hacker Phrases names
  hhc, hacker-phrases-colored               generate Colored Hacker Phrases names
  hhp, hacker-phrases-colored-piped         generate Colored Hacker Phrases names, let's work as a pipe
  i, internet                               generate Internet names
```

### Basic Usages

For example:

```bash
❯ ./bin/faker_darwin-amd64 a
    City                : East Eldred
    StreetName          : Lesley Locks
    StreetAddress       : 259 Donnelly Port
    SecondaryAddress    : Suite 586
    BuildingNumber      : 900
    BuildingNumber      : 929
    PostcodeByState     : 89678-3600
    ZipCode             : 35609
    ZipCodeByState      : 57406
    TimeZone            : Asia/Hong_Kong
    CityPrefix          : South
    CitySuffix          : mouth
    StreetSuffix        : Mills
    State               : Maryland
    StateAbbr           : GA
    Country             : Guatemala
    CountryCode         : MN
    Latitude            : -25.92453
    Longitude           : 62.222977
    String              : 418 Braun Roads Suite 343, Aprilbury Georgia 41333

❯ ./bin/faker_darwin-amd64 f --visa
        visa : 4364442188476

❯ ./bin/faker_darwin-amd64 f -m
    mastercard : 6771-8918-3284-3326

❯ ./bin/faker_darwin-amd64 c -10
      ISBN10 : 400189804-7

❯ ./bin/faker_darwin-amd64 c -h
...
Options:
  [Type]
  -a,   --abn                               generates a ABN code (default=false)
  -e13, --ean13                             generates a EAN13 code (default=false)
  -e8,  --ean8                              generates a EAN8 code (default=false)
  -10,  --isbn10                            generates a ISBN10 code (default=false)
  -13,  --isbn13                            generates a ISBN13 code (default=false)
  -r,   --rut                               generates a RUT code (default=false)
...
```

Following the documentation at [dmgk/faker](https://github.com/dmgk/faker) too.



### Installations & Usages

#### via Homebrew

As a macOS user, you may install `faker` via Homebrew:

```bash
brew install hedzr/brew/faker
```

Or,

```bash
brew tap hedzr/brew
brew install faker
```

#### via DockerHub or Github Docker Registry

Pull the docker image:
```bash
docker pull hedzr/faker              # from Docker Hub
docker pull ghcr.io/hedzr/cli/faker  # from Github Packages
```

Run as a local app:

```bash
docker run -it --rm hedzr/faker --help
docker run -it --rm hedzr/faker finance --mastercard
# or:
docker run -it --rm ghcr.io/hedzr/cli/faker
```

#### Else build from source codes



## Getting Started (for Developer)

To run the CLI app:

```bash
# go install -v github.com/swaggo/swag/cmd/swag
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


