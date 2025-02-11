 # Gofish - Redfish and Swordfish client library

[![Go Doc](https://godoc.org/github.com/ciferlu1024/gofish?status.svg)](http://godoc.org/github.com/ciferlu1024/gofish)
[![Go Report Card](https://goreportcard.com/badge/github.com/ciferlu1024/gofish?branch=main)](https://goreportcard.com/report/github.com/ciferlu1024/gofish)
[![Releases](https://img.shields.io/github/release/ciferlu1024/gofish/all.svg?style=flat-square)](https://github.com/ciferlu1024/gofish/releases)
[![LICENSE](https://img.shields.io/github/license/ciferlu1024/gofish.svg?style=flat-square)](https://github.com/ciferlu1024/gofish/blob/main/LICENSE)

![Gofish Logo](./images/gofish200x117.png)

## Introduction

Gofish is a Golang library for interacting with [DMTF
Redfish](https://www.dmtf.org/standards/redfish) and [SNIA
Swordfish](https://www.snia.org/forums/smi/swordfish) enabled devices.

## Usage ##

Basic usage would be:

```go

package main

import (
    "fmt"

    "github.com/ciferlu1024/gofish"
)

func main() {
    c, err := gofish.ConnectDefault("http://localhost:5000")
    if err != nil {
        panic(err)
    }

    service := c.Service
    chassis, err := service.Chassis()
    if err != nil {
        panic(err)
    }

    for _, chass := range chassis {
        fmt.Printf("Chassis: %#v\n\n", chass)
    }
}
```
