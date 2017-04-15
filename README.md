# openbd

[![Travis branch](https://img.shields.io/travis/osamingo/openbd/master.svg)](https://travis-ci.org/osamingo/openbd)
[![codecov](https://codecov.io/gh/osamingo/openbd/branch/master/graph/badge.svg)](https://codecov.io/gh/osamingo/openbd)
[![Go Report Card](https://goreportcard.com/badge/osamingo/openbd)](https://goreportcard.com/report/osamingo/openbd)
[![codebeat badge](https://codebeat.co/badges/b8889073-90a0-4c69-805a-aa8f69361236)](https://codebeat.co/projects/github-com-osamingo-openbd-master)
[![GoDoc](https://godoc.org/github.com/osamingo/openbd?status.svg)](https://godoc.org/github.com/osamingo/openbd)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/osamingo/openbd/master/LICENSE)

## About

- An [openbd](https://openbd.jp) client for Go.

## Install

```bash
$ go get -u github.com/osamingo/openbd
```

## Usage

```go
package main

import (
    "fmt"

    "github.com/osamingo/openbd"
)

func main() {

    cli, err := openbd.NewClientV1("https://api.openbd.jp", nil)
    if err != nil {
        panic(err)
    }
    
    isbn := "9784780802047"
    m, err := cli.Get(isbn)
    if err != nil {
        panic(err)
    }

    fmt.Println(m[isbn].Title())
}
```

## License

Released under the [MIT License](https://github.com/osamingo/openbd/blob/master/LICENSE).
