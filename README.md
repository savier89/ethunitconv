# ethunitconv
[![Build Status](https://travis-ci.org/savier89/ethunitconv.svg?branch=master)](https://travis-ci.org/savier89/ethunitconv)

Ethereum Wei unit Converter

Convert Ether units http://ethdocs.org/en/latest/ether.html

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
```