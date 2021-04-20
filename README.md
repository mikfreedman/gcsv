[![Go Report Card](https://goreportcard.com/badge/github.com/mikfreedman/gcsv)](https://goreportcard.com/report/github.com/mikfreedman/gcsv)
[![Go Build](https://github.com/mikfreedman/gcsv/actions/workflows/go.yml/badge.svg)](https://github.com/mikfreedman/gcsv/actions/workflows/go.yml/badge.svg)

Gomega matchers for CSV Schema
==================================

This package provides [Gomega](https://github.com/onsi/gomega) matchers to write assertions against Schema of a CSV:

RepresentSchema()
-------------------
Verifies that an entire CSV matches the actual schema of an `[]interface{}` array made up of the following basic types:

* `int`
* `bool`
* `float64`
* `string` 

*Note: Values for the actual schema are representative only and don't mean anything in and of themselves, they are just used for [switching on type](/match_schema.go#L57)

```go
import (
  . "github.com/mikfreedman/gcsv"
)

Expect("a,b,c,1,2,3").To(RepresentSchema([]interface{"a","b","c",1,2,3})) // Pass
Expect("a,b,c,d,e,f").To(RepresentSchema([]interface{"a","b","c",1,2,3})) // Fail!

Expect("header1,header2,header3,header4,header5,header6\na,b,c,d,e,f").To(RepresentSchema([]interface{"a","b","c",1,2,3}, IgnoreHeaderRow())) // Pass

```
