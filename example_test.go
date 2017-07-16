// Copyright 2017 The Bulma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bulma_test

import (
	"fmt"
	"net/http"

	"github.com/nanoninja/bulma"
)

func Example() {
	onSuccess := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Success")
	})

	ba := bulma.BasicAuth(bulma.Realm, onSuccess, bulma.User{
		"foo": "bar",
		"bar": "foo",
	})

	http.Handle("/admin", ba)
	http.ListenAndServe(":3000", nil)
}

func ExampleNew() {
	onSuccess := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Success")
	})

	onError := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Error")
	})

	validator := bulma.User{
		"foo": "bar",
		"bar": "foo",
	}

	ba := bulma.New(&bulma.Config{
		Realm:     "MyRealm",
		Validator: validator,
		Success:   onSuccess,
		Error:     onError,
	})

	http.Handle("/admin", ba)
	http.ListenAndServe(":3000", nil)
}

func ExampleValidateFunc() {
	onSuccess := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Success")
	})

	validator := bulma.ValidateFunc(func(c *bulma.Credential) bool {
		return c.Authorization && c.Username == "foo" && c.Password == "bar"
	})

	ba := bulma.BasicAuth(bulma.Realm, onSuccess, validator)

	http.Handle("/admin", ba)
	http.ListenAndServe(":3000", nil)
}

func ExampleAuth() {
	onSuccess := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Success")
	})

	ba := bulma.BasicAuth(bulma.Realm, onSuccess, bulma.Auth("foo", "bar"))

	http.Handle("/admin", ba)
	http.ListenAndServe(":3000", nil)
}
