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

The available commands are:

```bash
Commands:
  a, addr, address                          generate address
  app                                       generate app names
  av, avatar                                generate Avatar names
  btc, bitcoin                              generate Bitcoin (BTC) names
  bz, business, biz, credit-card            generate Business names
  c, code, codes                            generate Codes (ISBN10, ISBN13, EAN13, EAN8, RUT, ABN)
  cc, commerce                              generate Commerce names
  co, company, corp                         generate Company names
  f, finance, fin                           generate Finance names (visa, mastercard, ...)
  hh, hacker, hack                          generate Hacker names
  hp, hacker-phrases                        generate Hacker Phrases names
  hhc, hacker-phrases-colored               generate Colored Hacker Phrases names
  hhp, hacker-phrases-colored-piped         generate Colored Hacker Phrases names, let's work as a pipe
  i, internet                               generate Internet names
  l, lorem                                  generate Lorem text
  n, name                                   generate Name
  num, number                               generate address
  pn, phone-number                          generate PhoneNumber record
  t, team                                   generate Team record
```

### Basic Usages

For example:

```bash
❯ ./bin/faker_darwin-amd64 a --plain
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

❯ ./bin/faker_darwin-amd64 f --visa --plain # shortcut to 'finance'
        visa : 4364442188476

❯ ./bin/faker_darwin-amd64 finance -m
    mastercard : 6771-8918-3284-3326

❯ ./bin/faker_darwin-amd64 c -10 --plain    # shortcut to 'code'
      ISBN10 : 400189804-7

❯ ./bin/faker_darwin-amd64 code -h          # print options for 'code'
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

#### Output Formats

There are several output formats: `json`, `json-compact` (_default_), `yaml` or `plain`. You may specify its by `--json`, `--yaml` or `--plain`.

```bash
❯ ./bin/faker_darwin-amd64 c -10 --json            # shortcut to 'code'
{
  "Code": {
    "ISBN10": "821264042-6"
  }
}
```

> **Exceptions**:  
> 'hacker-phrases', 'hacker-phrases-colored' and 'hacker-phrases-colored-piped' have only plain formats:
> ```bash
> ❯ go run ./cli/app/cli/app hp
> If we parse the matrix, we can get to the CSS system through the bluetooth TCP interface!; We need to copy the auxiliary PCI bandwidth!; Try to program the COM driver, maybe it will reboot the digital bandwidth!; You can't compress the capacitor without indexing the optical USB driver!; Use the cross-platform TCP alarm, then you can generate the digital system!; The SQL interface is down, override the haptic protocol so we can navigate the XSS protocol!; Bypassing the matrix won't do anything, we need to synthesize the bluetooth RSS driver!; I'll generate the multi-byte SSL card, that should array the JSON panel!; If we index the pixel, we can get to the JSON application through the auxiliary JBOD bandwidth!; We need to back up the auxiliary TCP monitor!; Try to parse the JSON pixel, maybe it will override the 1080p application!; You can't connect the system without backing up the solid state USB protocol!; Use the back-end SMTP firewall, then you can parse the digital feed!; The HDD interface is down, compress the wireless sensor so we can synthesize the XSS system!; Indexing the program won't do anything, we need to parse the online EXE firewall!; I'll index the optical IB circuit, that should array the JSON interface!; If we program the transmitter, we can get to the ADP hard drive through the virtual JSON bandwidth!; We need to bypass the auxiliary CSS firewall!; Try to program the ADP pixel, maybe it will index the mobile alarm!; You can't transmit the matrix without programming the digital XML card!; Use the cross-platform COM array, then you can override the cross-platform bus!; The FTP panel is down, copy the virtual application so we can quantify the FTP feed!; Copying the driver won't do anything, we need to compress the cross-platform JBOD matrix!; I'll parse the digital SSL hard drive, that should hard drive the PNG card!; If we copy the card, we can get to the SMS card through the 1080p CSS feed!; We need to compress the open-source XSS card!; Try to copy the XSS card, maybe it will generate the auxiliary array!; You can't connect the monitor without programming the multi-byte SSL pixel!; Use the 1080p SCSI port, then you can generate the solid state bandwidth!; The SMTP protocol is down, generate the neural transmitter so we can input the SDD alarm!; Navigating the program won't do anything, we need to bypass the cross-platform IB feed!; I'll compress the open-source GB array, that should circuit the SCSI hard drive!
> ```

#### Output one field alone or Dump All

You may print just one field in a record. For instance, the following command print Title field of a Name record:

```bash
❯ go run ./cli/app/cli/app --json name -t
{
  "Name": {
    "Title": "Lead Data Technician"
  }
}
```

In another side, all fields can be printed of course:

```bash
❯ go run ./cli/app/cli/app --json n
{
  "Name": {
    "FirstName": "Fritz",
    "LastName": "Kilback",
    "Name": "Kenneth Becker",
    "Prefix": "Ms.",
    "String": "Sammie Windler",
    "Suffix": "Sr.",
    "Title": "Principal Configuration Agent"
  }
}
```

Check out the options of `Name` record:

```bash
❯ go run ./cli/app/cli/app --json n -h
...
Options:
  [Type]
  -f,  --first-name,--fn                    generates a firstName field (default=false)
  -l,  --last-name,--ln                     generates a lastName field (default=false)
  -n,  --name                               generates a fullName field (default=false)
  -p,  --prefix                             generates a prefix field (default=false)
  -ss, --string,--str,--sz                  generates a string field (default=false)
  -s,  --suffix                             generates a suffix field (default=false)
  -t,  --title                              generates a title field (default=false)
...
```

#### Which supports `Output one field alone or all`

Totally, these records have the feature:

- App
- Address
- Avatar
- Bitcoin
- Business
- Commerce
- Company
- Internet
- Lorem
- Name
- Number
- PhoneNumber



## Installations & Usages

### Homebrew

As a macOS user, you may install `faker` via Homebrew:

```bash
brew install hedzr/brew/faker
```

Or,

```bash
brew tap hedzr/brew
brew install faker
```

Now, `faker` is available.


### DockerHub or Github Docker Registry

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

### Else build from source codes



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


