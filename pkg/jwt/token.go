// Copyright 2017. All rights reserved.
// Use of this source code is governed by a BSD 3-Clause License
// license that can be found in the LICENSE file.

package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"strings"
)

// Token is a struct that maps all parts of a JWT Token.
// Different fields will be used depending on whether you're
// creating or parsing/verifying a token.
type Token struct {
	Raw       string
	Header    map[string]string
	Claims    Claims
	Signature string
	Valid     bool
}

// NewToken creates a new token with the specified id and expiration delta
func (jwt *JWT) NewToken(id string, plusExpire int64) Token {
	return Token{
		Header: map[string]string{
			"typ": "JWT",
			"alg": jwt.AlgName,
		},
		Claims: jwt.newClaims(id, plusExpire),
	}
}

// SignedString returns the complete, signed token
func (jwt *JWT) SignedString(t Token) (string, error) {
	var sig, sstr string
	var err error
	if sstr, err = t.signingString(); err != nil {
		return "", err
	}
	if sig, err = jwt.sign(sstr); err != nil {
		return "", err
	}
	return strings.Join([]string{sstr, sig}, "."), nil
}

func (jwt *JWT) sign(signingString string) (string, error) {
	hasher := hmac.New(sha256.New, jwt.Key)
	_, err := hasher.Write([]byte(signingString))
	if err != nil {
		return "", err
	}
	return encodeSegment(hasher.Sum(nil)), nil
}

func (t *Token) signingString() (string, error) {
	var err error
	parts := make([]string, 2)

	var jsonValue []byte

	if jsonValue, err = json.Marshal(t.Header); err != nil {
		return "", err
	}
	parts[0] = encodeSegment(jsonValue)

	if jsonValue, err = json.Marshal(t.Claims); err != nil {
		return "", err
	}
	parts[1] = encodeSegment(jsonValue)

	return strings.Join(parts, "."), nil
}
