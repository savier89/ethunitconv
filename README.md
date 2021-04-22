# ethunitconv
[![Build Status](https://travis-ci.org/savier89/ethunitconv.svg?branch=master)](https://travis-ci.org/savier89/ethunitconv)
[![Go Report Card](https://goreportcard.com/badge/github.com/savier89/ethunitconv)](https://goreportcard.com/report/github.com/savier89/ethunitconv)

Ethereum Wei unit Converter

## Features

Convert all Ether units http://ethdocs.org/en/latest/ether.html

## Install

    go get -v github.com/savier89/ethunitconv


## Usage

```
package main

import (
	"fmt"

	"github.com/savier89/ethunitconv"
)

func main() {
	ether := ethunitconv.FromWei("418160973408927260000000000", "Ether")
	fmt.Println(ether)
	ether2wei := ethunitconv.ToWei(ether, "Ether")
	fmt.Println(ether2wei)
}
```

## Output
```
$ go run main.go
418160973.4089272600
418160973408927259994095616

$ go test -v
=== RUN   TestConverter
418160973.4089272600
418160973408927259994095616
--- PASS: TestConverter (0.00s)
PASS
ok      ~/go-projects/ethunitconv        0.350s

```

## Donate

You can help fund the developer by donating to the following wallet addresses:

	BTC: 1HhcW3NjAswZPGBmHo1Kcr2JPA3cfn1bA6
	ETH: 0xF9E93f163E8Bd606BD36CA2D177C345750C66241
	CLO: 0x9c0B6195092780963ef2EB978C4b2E1d2f375Bb4
	
