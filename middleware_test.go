// Copyright 2016 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bulma

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMiddleware(t *testing.T) {
	for _, tt := range getBasicAuthHandlerTests {
		mw := Middleware(credentialTest.username, credentialTest.password, DefaultRealm)
		handler := mw(handlerTest)

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

func TestMiddlewareFunc(t *testing.T) {
	for _, tt := range getBasicAuthHandlerTests {
		ba := BasicAuth(credentialTest.username, credentialTest.password)
		mw := MiddlewareFunc(ba, DefaultRealm)

		handler := mw(handlerTest)
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
