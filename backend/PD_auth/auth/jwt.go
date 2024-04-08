package auth

import (
	"fmt"
	"net/http"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("aodjfeowjedsoj134trefdsas234t")

func TokenVerificationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token JWT ausente na requisição")
			return
		}

		claims := &jwt.StandardClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token JWT inválido")
			return
		}

		next.ServeHTTP(w, r)
	})
}
