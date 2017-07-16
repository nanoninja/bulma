// Copyright 2017 The Bulma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package bulma

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var validatorTest = []struct {
	validator Validator
	code      int
	body      string
}{
	{
		Auth(userTest.username, userTest.password),
		http.StatusOK,
		"success",
	},
	{
		User{userTest.username: userTest.password},
		http.StatusOK,
		"success",
	},
}

func TestValidator(t *testing.T) {
	for _, tt := range validatorTest {
		basicAuth := BasicAuth(Realm, handlerTest, tt.validator)

		rec := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "http://example.com", nil)
		if err != nil {
			t.Error(err)
		}

		req.SetBasicAuth(userTest.username, userTest.password)
		basicAuth.ServeHTTP(rec, req)

		if rec.Code != tt.code {
			t.Errorf("got status %d; want %d", rec.Code, http.StatusOK)
		}
		if got, want := rec.Body.String(), tt.body; got != want {
			t.Errorf("rec.Body.String() got %q; want %q", got, want)
		}

		username, password, ok := req.BasicAuth()
		if !ok && username != userTest.username && password != userTest.password {
			t.Errorf(
				"req.BasicAuth() got (%q/%q); want (%q/%q)",
				username,
				password,
				userTest.username,
				userTest.password,
			)
		}

		req.Header.Del("Authorization")
		rec = httptest.NewRecorder()
		basicAuth.ServeHTTP(rec, req)

		if rec.Code != http.StatusUnauthorized {
			t.Errorf("got status %d; want %d", rec.Code, http.StatusOK)
		}
		if got, want := rec.Body.String(), "Unauthorized"; got != want {
			t.Errorf("rec.Body.String() got %q; want %q", got, want)
		}
	}
}
