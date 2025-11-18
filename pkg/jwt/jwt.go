package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Claims describes the standard claims we encode in tokens.
type Claims struct {
	UserID string   `json:"uid"`
	Roles  []string `json:"roles"`
	jwt.StandardClaims
}

// SignToken creates a signed JWT using the supplied secret key.
func SignToken(secret string, userID string, roles []string, ttl time.Duration) (string, error) {
	claims := Claims{
		UserID: userID,
		Roles:  roles,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ParseToken validates and extracts claims from a token string.
func ParseToken(secret, tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrSignatureInvalid
}
