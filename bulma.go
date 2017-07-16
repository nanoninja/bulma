// Copyright 2017 The Bulma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bulma provides HTTP Basic Auth.
package bulma

import (
	"net/http"
)

// Realm is the authenticate message to require authorization.
// See RFC 2617, section 1.2.
var Realm = "Authorization Required"

// Config is used to configure the basic auth handler.
type Config struct {
	// Realm Default is "Authorization Required"
	// See RFC 2617, section 1.2.
	Realm string

	// Validator is the interface indicating the type implementing it
	// supports credentials validation.
	Validator Validator

	// Success will be used if the authentication succeeded.
	Success http.Handler

	// Error will be used if authentication fails.
	Error http.Handler
}

type basicAuth struct {
	config *Config
}

// New returns an http handler from a given configuration.
func New(c *Config) http.Handler {
	if c.Realm == "" {
		c.Realm = Realm
	}
	return &basicAuth{c}
}

func (b *basicAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	credential := Credential{username, password, ok}

	var next http.Handler

	if b.config.Validator != nil && b.config.Validator.Validate(&credential) {
		next = b.config.Success
	} else if b.config.Error != nil {
		next = err(b.config.Realm, b.config.Error)
	} else {
		next = require(b.config.Realm)
	}

	next.ServeHTTP(w, r)
}

// BasicAuth is a simple wrapper of New.
func BasicAuth(realm string, success http.Handler, v Validator) http.Handler {
	return New(&Config{
		Realm:     realm,
		Success:   success,
		Validator: v,
	})
}

// require returns a Handler that authenticates user.
func require(realm string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeHeader(w, realm)
		w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
	})
}

// writeHeader replies to the request with the authentication header.
func writeHeader(w http.ResponseWriter, realm string) {
	w.Header().Add("WWW-Authenticate", `Basic realm="`+realm+`"`)
	w.WriteHeader(http.StatusUnauthorized)
}

func err(realm string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeHeader(w, realm)
		h.ServeHTTP(w, r)
	})
}
