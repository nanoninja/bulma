// Copyright 2017 The Bulma Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bulma

var (
	_ Validator = (*ValidateFunc)(nil)
	_ Validator = (*User)(nil)
)

// Credential is provided when the client is requesting access to
// protected resources.
type Credential struct {
	Username      string
	Password      string
	Authorization bool
}

// Validator is the interface indicating the type implementing it
// supports credentials validation.
type Validator interface {
	// Validate validates the given credentials and
	// returns false if validation fails.
	Validate(*Credential) bool
}

// ValidateFunc implements a validator.
type ValidateFunc func(*Credential) bool

// Validate validates the given credentials and
// returns false if validation fails.
func (f ValidateFunc) Validate(c *Credential) bool {
	return f(c)
}

// Auth is a basic credentials function implementing a validator.
func Auth(username, password string) Validator {
	return ValidateFunc(func(c *Credential) bool {
		return c.Authorization && c.Username == username && c.Password == password
	})
}

// User is a list of users implementing a validator.
type User map[string]string

// Validate validates the given credentials and
// returns false if validation fails.
func (u User) Validate(c *Credential) bool {
	password, ok := u[c.Username]
	return ok && Auth(c.Username, password).Validate(c)
}
