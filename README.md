[![Go Report Card](https://goreportcard.com/badge/github.com/mikfreedman/gcsv)](https://goreportcard.com/report/github.com/mikfreedman/gcsv)

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


```go
import (
  . "github.com/mikfreedman/gcsv"
)

Expect("a,b,c,1,2,3").To(RepresentSchema([]interace{"a","b","c",1,2,3})) // Pass
Expect("a,b,c,d,e,f").To(RepresentSchema([]interace{"a","b","c",1,2,3})) // Fail!

Expect("header1,header2,header3,header4,header5,header6\na,b,c,d,e,f").To(RepresentSchema([]interace{"a","b","c",1,2,3}, IgnoreHeaderRow())) // Pass

```
