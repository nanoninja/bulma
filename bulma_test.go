// Copyright 2016 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bulma

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type basicAuthCredentialsTest struct {
	username, password string
}

var getBasicAuthHandlerTests = []struct {
	username string
	password string
	code     int
}{
	{"", "", 401},
	{"foo", "bar", 401},
	{"username", "password", 200},
	{"x", "y", 401},
	{"à", "é", 401},
	{"tofoo", "gobar", 401},
	{"abcd", "1234", 401},
}

var handlerTest = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Open sesame")
})

var credentialTest = basicAuthCredentialsTest{"username", "password"}

func TestDefaultRealm(t *testing.T) {
	expected := "Authorization Required"
	if DefaultRealm != expected {
		t.Errorf("got DefaultRealm = %s, want %s", DefaultRealm, expected)
	}
}

func TestBasicAuthHandler(t *testing.T) {
	for _, tt := range getBasicAuthHandlerTests {
		auth := BasicAuth(credentialTest.username, credentialTest.password)
		handler := BasicAuthHandler(auth, handlerTest, DefaultRealm)
		req, _ := http.NewRequest("GET", "http://example.com/", nil)
		req.SetBasicAuth(tt.username, tt.password)

		rw := httptest.NewRecorder()
		handler.ServeHTTP(rw, req)

		user, pass, _ := req.BasicAuth()
		if user != tt.username || pass != tt.password {
			t.Errorf("got user = %v and pass = %v, want %s and %s", user, pass, tt.username, tt.password)
		}

		if tt.code != rw.Code {
			t.Errorf("got code = %d, want %d", rw.Code, tt.code)
		}
	}
}
