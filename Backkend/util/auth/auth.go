package auth

import (
	"fmt"
	"net/http"
	"wishticket/util/jwt"
)

func GetJWTTokenFromHeader(r *http.Request) (string, error) {
	jwtString := r.Header.Get("auth")
	if jwtString == "" {
		return "", fmt.Errorf("missing authorization header")
	}
	return jwtString, nil
}


// GetJWTPayloadFromHeader extracts the JWT payload from the Authorization header of an HTTP request.
// It first retrieves the JWT token from the header, verifies the token, and then decodes the payload.
//
// Parameters:
//   r (*http.Request): The HTTP request containing the Authorization header with the JWT token.
//
// Returns:
//   (jwt.JWTPayload, error): Returns the decoded JWT payload if successful, otherwise returns an error.
func GetJWTPayloadFromHeader(r *http.Request) (jwt.JWTPayload, error) {
	jwtToken, err := GetJWTTokenFromHeader(r)
	var jwtData jwt.JWTPayload
	if err != nil {
		return jwtData, err
	}
	err = jwt.VerifyToken(jwtToken)
	if err != nil {
		return jwtData, err
	}
	jwtData, err = jwt.DecodeBearer(jwtToken)
	if err != nil {
		return jwtData, err
	}
	return jwtData, err
}
