package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"strings"
)

// ParseString allocates and returns a new Token based on a JWT signed token string
func (jwt *JWT) ParseString(tokenString string) (*Token, error) {
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return nil, ErrInvalidNumberOfSegments
	}

	var err error
	token := &Token{Raw: tokenString}

	var headerBytes []byte
	if headerBytes, err = decodeSegment(parts[0]); err != nil {
		return token, err
	}

	if err = json.Unmarshal(headerBytes, &token.Header); err != nil {
		return token, ErrMalformedTokenHeader
	}

	var claimBytes []byte
	if claimBytes, err = decodeSegment(parts[1]); err != nil {
		return token, err
	}

	if err = json.Unmarshal(claimBytes, &token.Claims); err != nil {
		return token, ErrMalformedTokenClaims
	}

	if token.Header["alg"] != jwt.AlgName {
		return token, ErrInvalidAlgorithm
	}

	if err = jwt.validateClaims(token.Claims); err != nil {
		return token, err
	}

	token.Signature = parts[2]
	if err = jwt.verify(strings.Join(parts[0:2], "."), token.Signature); err != nil {
		return token, err
	}

	token.Valid = true

	return token, nil
}

func (jwt *JWT) verify(signingString, signature string) error {
	sig, err := decodeSegment(signature)
	if err != nil {
		return err
	}
	hasher := hmac.New(sha256.New, jwt.Key)
	hasher.Write([]byte(signingString))
	if !hmac.Equal(sig, hasher.Sum(nil)) {
		return ErrInvalidSignature
	}
	return nil
}
