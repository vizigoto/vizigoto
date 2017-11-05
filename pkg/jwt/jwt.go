// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package jwt

// JWT contains the main api
type JWT struct {
	Audience string
	Issuer   string
	Subject  string
	AlgName  string
	Key      []byte
}

var algName = "HS256"

// New allocates and returns a new JWT.
func New(audience string, issuer string, subject string, key string) *JWT {
	return &JWT{audience, issuer, subject, algName, []byte(key)}
}
