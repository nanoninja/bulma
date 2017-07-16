# Bulma

Basic authentication implementation for Go.

[![license](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://github.com/nanoninja/bulma/blob/master/LICENSE) [![godoc](https://godoc.org/github.com/nanoninja/bulma?status.svg)](https://godoc.org/github.com/nanoninja/bulma)
[![Go project version](https://badge.fury.io/go/github.com%2Fnanoninja%2Fbulma.svg)](https://badge.fury.io/go/github.com%2Fnanoninja%2Fbulma)
[![build Status](https://travis-ci.org/nanoninja/bulma.svg)](https://travis-ci.org/nanoninja/bulma)
[![Coverage Status](https://coveralls.io/repos/github/nanoninja/bulma/badge.svg?branch=master)](https://coveralls.io/github/nanoninja/bulma?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/nanoninja/bulma)](https://goreportcard.com/report/github.com/nanoninja/bulma) [![codebeat](https://codebeat.co/badges/58e89ce4-2fd8-4a93-b624-afdbbb44a6e3)](https://codebeat.co/projects/github-com-nanoninja-bulma)

## Installation

```sh
go get github.com/nanoninja/bulma
```

## Getting started
After installing Go and setting up your
[GOPATH](http://golang.org/doc/code.html#GOPATH), create your first `.go` file.

```go
package main

import (
    "fmt"
    "net/http"

    "github.com/nanoninja/bulma"
)

func main() {
    onSuccess := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Dashboard")
    })

    ba := bulma.BasicAuth(bulma.Realm, onSuccess, bulma.User{
        "foo": "bar",
        "bar": "foo",
    })

    http.Handle("/admin", ba)
    http.ListenAndServe(":3000", nil)
}
```

[Open your favorite browser](http://localhost:3000/admin)

## Using a function as validator

```go
func main() {
    onSuccess := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Dashboard")
    })

    validator := bulma.ValidateFunc(func(c *bulma.Credential) bool {
        return c.Authorization && c.Username == "foo" && c.Password == "bar"
    })

    ba := bulma.BasicAuth(bulma.Realm, onSuccess, validator)

    http.Handle("/admin", ba)
    http.ListenAndServe(":3000", nil)
}
```

## Using configuration
The configuration allows you to set up HTTP authentication.

```go
type Config struct {
    Realm string
    Validator Validator
    Success http.Handler
    Error http.Handler
}
```

Example :

```go
func main() {
    onSuccess := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Dashboard")
    })

    onError := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "My Custom Error Handler")
    })

    ba := bulma.New(&bulma.Config{
        Realm:     "MyRealm",
        Validator: bulma.Auth("foo", "bar"),
        Success:   onSuccess,
        Error:     onError,
    })

    http.Handle("/admin", ba)
    http.ListenAndServe(":3000", nil)
}
```

## Creating Validator
To create a validator, use bulma.Validator interface.

```go
type Validator interface {
    Validate(*Credential) bool
}
```

Example :

```go
type MyValidator struct {
    username, password string
}

func (v MyValidator) Validate(c *bulma.Credential) bool {
    return c.Authorization && v.username == c.Username && v.password == c.Password
}
```

Using your own validator

```go
func main() {
    onSuccess := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Dashboard")
    })

    ba := bulma.BasicAuth("MyRealm", onSuccess, MyValidator{"foo", "bar"})

    http.Handle("/admin", ba)
    http.ListenAndServe(":3000", nil)
}
```

## License

Bulma is licensed under the Creative Commons Attribution 3.0 License, and code is licensed under a [BSD license](https://github.com/nanoninja/bulma/blob/master/LICENSE).