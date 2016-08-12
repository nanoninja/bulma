// Copyright 2016 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bulma provides HTTP Basic Auth.
package bulma

import "net/http"

// DefaultRealm is the authentication message by default.
const DefaultRealm = "Authorization Required"

// Authentifier interface is used by basic HTTP authentication handler.
type Authentifier interface {
	// Authenticate checks the permissions to access the website.
	Authenticate(r *http.Request) bool

	// RequireAuth asks user credentials.
	RequireAuth(rw http.ResponseWriter, realm string)
}

// BasicAuthFunc represents the authentication credentials.
type BasicAuthFunc func(username, password string) bool

// BasicAuth is a helper that wraps
// a call to a function returning BasicAuthFunc.
func BasicAuth(username, password string) BasicAuthFunc {
	return func(user, pass string) bool {
		return username == user && password == pass
	}
}

// RequireAuth sets the response Authorization header to use HTTP
// Basic Authentication with the provided username and password.
//
// With HTTP Basic Authentication the provided
// username and password are not encrypted.
func (f BasicAuthFunc) RequireAuth(rw http.ResponseWriter, realm string) {
	rw.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
	rw.WriteHeader(401)
}

// Authenticate returns true if and only if
// the authentication credentials are valid.
func (f BasicAuthFunc) Authenticate(r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	return ok && f(username, password)
}

type basicAuthHandler struct {
	auth  Authentifier
	realm string
	next  http.Handler
}

// BasicAuthHandler returns a handler that serves basic authentication HTTP.
//
// The typical use case for BasicAuthHandler is to register a
// BasicAuthFunc function.
//
//    package main
//
//    import (
//        "net/http"
//
//        "github.com/nanoninja/bulma"
//    )
//
//    func main() {
//        auth := bulma.BasicAuth("foo", "bar")
//
//        next := http.Handlerfunc(rw http.ResponseWriter, r *http.Request) {
//            rw.Write([]byte("Open sesame"))
//        })
//
//        handler := bulma.BasicAuthHandler(auth, next)
//
//        http.Handle("/sesame", handler)
//        http.ListenAndServe(":3000", nil)
//    }
func BasicAuthHandler(a Authentifier, h http.Handler, realm string) http.Handler {
	return &basicAuthHandler{
		auth:  a,
		next:  h,
		realm: realm,
	}
}

func (h *basicAuthHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if !h.auth.Authenticate(r) {
		h.auth.RequireAuth(rw, h.realm)
	} else {
		h.next.ServeHTTP(rw, r)
	}
}
