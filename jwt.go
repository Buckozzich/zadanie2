package app

import (
	"fmt"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"os"
	"ctx"
)

var JwtAuthentication struct

func JwtAuthentication(http.Handler) http.Handler  {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
		NoAout := []string{"/api/user/new", "api/user/login"}
	requestPath := r.URL.Path
for _, v := range  NoAout {
	if v == requestPath {
		serverHTTP(w,r)
		return
	}


	}	

}

	)
}


response := make(map[string]interface{})
tokenHeader := r.Header.Ger("Authorization")

if tokenHeader == ""

{response := u.Messenge(false, "Not Working auth token")
w.Header().Add("Context-Type", "application/json")
u.Responde(w, response)
retrun}

splitted := strings.Split(tokenHeader, " ") 
		if len(splitted) != 2 {
			response = u.Message(false, "Invalid auth token")
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

	tokenPart := splitted[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []string(os.Getenv("token_password")), nil
		})	

		if err != nil {
			response = u.Message(false, "Malformed authentication token")
		 w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		if !token.Valid{ 
			response = u.Message(false, "Token is not valid.")
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) 

		
