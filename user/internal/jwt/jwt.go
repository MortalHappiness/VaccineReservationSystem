package jwt

import "github.com/golang-jwt/jwt/v4"

// Claims is a custom claims struct for JWT tokens.
type Claims struct {
	NationID string `json:"nationID"`
	jwt.StandardClaims
}

// NewToken creates a new JWT token.
func NewToken(nationID string, secret string) (string, error) {
	claims := Claims{
		NationID: nationID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ParseToken parses a JWT token.
func ParseToken(tokenString string, secret string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
