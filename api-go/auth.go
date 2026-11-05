package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func login(w http.ResponseWriter, r *http.Request) {
	if permitirNavegador(w, r) {
		return
	}
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var body struct {
		Usuario    string `json:"usuario"`
		Contrasena string `json:"contrasena"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if body.Usuario != os.Getenv("API_USER") || body.Contrasena != os.Getenv("API_PASSWORD") {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "credenciales invalidas"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   body.Usuario,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
	})
	firmado, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": firmado})
}

func requiereToken(siguiente http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if permitirNavegador(w, r) {
			return
		}
		encabezado := r.Header.Get("Authorization")
		if !strings.HasPrefix(encabezado, "Bearer ") {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "token requerido"})
			return
		}
		texto := strings.TrimPrefix(encabezado, "Bearer ")
		token, err := jwt.Parse(texto, func(t *jwt.Token) (any, error) {
			if t.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("firma invalida")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "token invalido"})
			return
		}
		siguiente(w, r)
	}
}