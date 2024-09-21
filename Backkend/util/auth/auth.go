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
