# parsestr

[![Build Status](https://travis-ci.org/donatj/parsestr.svg?branch=master)](https://travis-ci.org/donatj/parsestr)
[![GoDoc](https://godoc.org/github.com/donatj/parsestr?status.svg)](https://godoc.org/github.com/donatj/parsestr)

Parsestr is a simple Golang library emulating the query parsing behaviour of the PHP [parse_str](http://php.net/manual/en/function.parse-str.php).

Parsestr works by extending the Golang standard url.Values allowing parsing of PHP style `foo[bar]=baz&qux[]=quux` query strings.
