# parsestr

[![GoDoc](https://godoc.org/github.com/donatj/parsestr?status.svg)](https://godoc.org/github.com/donatj/parsestr)
[![CI](https://github.com/donatj/parsestr/actions/workflows/ci.yml/badge.svg)](https://github.com/donatj/parsestr/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/donatj/parsestr)](https://goreportcard.com/report/github.com/donatj/parsestr)

Parsestr is a simple Golang library emulating the query parsing behavior of the PHP [parse_str](http://php.net/manual/en/function.parse-str.php).

Parsestr works by extending the Golang standard url.Values allowing parsing of PHP style `foo[bar]=baz&qux[]=quux` query strings.

## Example

```go
package main

import (
	"fmt"
	"github.com/donatj/parsestr"
)

func main() {
	v, err := parsestr.ParseQuery("foo[bar]=baz&foo[qux]=quux&dog[]=cat")
	if err != nil {
		panic(err)
	}

	foo := v.GetSub("foo")
	fmt.Println(foo.Get("bar")) // baz
	fmt.Println(foo.Get("qux")) // quux

	dog := v.GetSub("dog")
	fmt.Println(dog.Get("")) // cat
}
```
