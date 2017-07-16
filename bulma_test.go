// Copyright 2017 The Bulma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package bulma

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type credentialTest struct {
	ok       bool
	username string
	password string
}

var basicAuthTest = []struct {
	username string
	password string
	code     int
}{
	{"Aladdin", "open sesame", http.StatusOK},
	{"", "", http.StatusUnauthorized},
	{"Foo", "bar", http.StatusUnauthorized},
	{"abc", "xyz", http.StatusUnauthorized},
	{"à", "é", http.StatusUnauthorized},
	{"Gopher", "golang", http.StatusUnauthorized},
	{"abcde", "12345", http.StatusUnauthorized},
}

var handlerTest = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("success"))
})

var userTest = credentialTest{false, "Aladdin", "open sesame"}

func testBasicAuth(basicAuth http.Handler, t *testing.T) {
	for _, tt := range basicAuthTest {
		rec := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "http://example.com/", nil)
		if err != nil {
			t.Error(err)
		}

		req.SetBasicAuth(tt.username, tt.password)
		basicAuth.ServeHTTP(rec, req)
		username, password, ok := req.BasicAuth()

		if !ok {
			t.Errorf("got authorization %v; want %v", ok, true)
		}
		if tt.username != username && tt.password != password {
			t.Errorf(
				"got username/password (%q/%q); want (%q/%q)",
				username,
				password,
				tt.username,
				tt.password,
			)
		}
		if tt.code != rec.Code {
			t.Errorf("got status code %d; want %d", tt.code, rec.Code)
		}
	}
}

func TestBasicAuth(t *testing.T) {
	var handlers = []struct {
		handler http.Handler
	}{
		{
			New(&Config{
				Validator: Auth(userTest.username, userTest.password),
				Success:   handlerTest,
				Error: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
				}),
			}),
		},
		{
			BasicAuth(Realm, handlerTest, User{
				userTest.username: userTest.password,
			}),
		},
	}
	for _, tt := range handlers {
		testBasicAuth(tt.handler, t)
	}
}
