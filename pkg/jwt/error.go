package jwt

import "errors"

// ErrTokenExpired is used when the token has a claims.ExpiresAt < now
var ErrTokenExpired = errors.New("jwt: token expired")

// ErrTokenUsedBeforeIssued is returned when the client try to use a token with now <= claims.IssuedAt >= now.
var ErrTokenUsedBeforeIssued = errors.New("jwt: token used before issued")

// ErrTokenNotValidYet is returned when the client try to use a token before the claims.NotBefore time
var ErrTokenNotValidYet = errors.New("jwt: token is not valid yet")

// ErrInvalidIssuer is returned when the claims.Issuer doesn't match
var ErrInvalidIssuer = errors.New("jwt: invalid issuer")

// ErrInvalidSubject is returned when the claims.Subject doesn't match
var ErrInvalidSubject = errors.New("jwt: invalid subject")

// ErrInvalidAudience is returned when the claims.Audience doesn't match
var ErrInvalidAudience = errors.New("jwt: invalid audience")

// ErrInvalidAlgorithm is returned when the token header key alg doesn't match
var ErrInvalidAlgorithm = errors.New("jwt: invalid algorithm")

// ErrInvalidSignature is returned when the token signature doesn't match
var ErrInvalidSignature = errors.New("jwt: invalid signature")

// ErrInvalidNumberOfSegments is returned when the token has less than 3 parts
var ErrInvalidNumberOfSegments = errors.New("jwt: token contains an invalid number of segments")

// ErrMalformedTokenHeader is returned when a error ocurrend with the jsaon Unmarshalling from token head segment
var ErrMalformedTokenHeader = errors.New("jwt: malformed token header string")

// ErrMalformedTokenClaims is returned when a error ocurrend with the jsaon Unmarshalling from token claims segment
var ErrMalformedTokenClaims = errors.New("jwt: malformed token claims string")
