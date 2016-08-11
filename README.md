# Bulma    [![Build Status](https://travis-ci.org/nanoninja/bulma.svg)](https://travis-ci.org/nanoninja/bulma) [![coverage](https://img.shields.io/badge/coverage-100%-brightgreen.svg?style=flat)](http://gocover.io/github.com/nanoninja/bulma) [![codebeat badge](https://codebeat.co/badges/58e89ce4-2fd8-4a93-b624-afdbbb44a6e3)](https://codebeat.co/projects/github-com-nanoninja-bulma) [![Go Report Card](https://goreportcard.com/badge/github.com/nanoninja/bulma)](https://goreportcard.com/report/github.com/nanoninja/bulma) [![go-doc](https://godoc.org/github.com/nanoninja/bulma?status.svg)](https://godoc.org/github.com/nanoninja/bulma)

Basic authentication implementation for Go.

## Installation

    go get github.com/nanoninja/bulma

## Getting started

After installing Go and setting up your
[GOPATH](http://golang.org/doc/code.html#GOPATH), create your first `.go` file.

``` go
package main

import (
	"fmt"
	"net/http"

	"github.com/nanoninja/bulma"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Welcome to the home page!")
	})

    ba := bulma.BasicAuth("username", "password")
	handler := bulma.BasicAuthHandler(ba, mux, bulma.DefaultRealm)

	http.ListenAndServe(":3000", handler)
}
```

## Using callback

Enfore basic authentication by providing a BasicAuthFunc,
which must return true in order to gain access.
``` go
package main

import (
	"fmt"
	"net/http"

	"github.com/nanoninja/bulma"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Welcome to the home page!")
	})

    f := func(user, pass string) bool {
        return user == "username" && pass == "password"
    }

	handler := bulma.BasicAuthHandler(f, mux, bulma.DefaultRealm)

	http.ListenAndServe(":3000", handler)
}
```

## License

Bulma is licensed under the Creative Commons Attribution 3.0 License, and code is licensed under a [BSD license](https://github.com/nanoninja/bulma/blob/master/LICENSE).
