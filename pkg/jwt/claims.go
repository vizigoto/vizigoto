package jwt

import (
	"crypto/subtle"
	"time"
)

// Claims is a structured version of Claims Section, as referenced at
// https://tools.ietf.org/html/rfc7519#section-4.1
type Claims struct {
	Audience  string `json:"aud"`
	ExpiresAt int64  `json:"exp"`
	ID        string `json:"jti"`
	IssuedAt  int64  `json:"iat"`
	Issuer    string `json:"iss"`
	NotBefore int64  `json:"nbf"`
	Subject   string `json:"sub"`
}

func (jwt *JWT) newClaims(id string, plusExpire int64) Claims {
	expire := time.Now().Unix() + plusExpire
	return Claims{
		Audience:  jwt.Audience,
		ExpiresAt: expire,
		ID:        id,
		IssuedAt:  time.Now().Unix(),
		Issuer:    jwt.Issuer,
		NotBefore: time.Now().Unix(),
		Subject:   jwt.Subject,
	}
}

func (jwt *JWT) validateClaims(c Claims) error {
	now := time.Now().Unix()
	if subtle.ConstantTimeCompare([]byte(c.Audience), []byte(jwt.Audience)) == 0 {
		return ErrInvalidAudience
	}
	if now >= c.ExpiresAt {
		return ErrTokenExpired
	}
	if now <= c.IssuedAt {
		return ErrTokenUsedBeforeIssued
	}
	if subtle.ConstantTimeCompare([]byte(c.Issuer), []byte(jwt.Issuer)) == 0 {
		return ErrInvalidIssuer
	}
	if now <= c.NotBefore {
		return ErrTokenNotValidYet
	}
	if subtle.ConstantTimeCompare([]byte(c.Subject), []byte(jwt.Subject)) == 0 {
		return ErrInvalidSubject
	}
	return nil
}
